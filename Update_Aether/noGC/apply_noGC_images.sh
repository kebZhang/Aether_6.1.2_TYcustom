#!/bin/bash
# Run on machine B (K8s node). Same as ../apply_aether_images.sh, for the noGC
# image. Only switches the images; GC is set by set_aether_noGC.sh, which must
# run AFTER this: helm upgrade puts amf's GOGC=200 and smf's command back.
set -euo pipefail

TAG=NF8HTTP_noGC_1109                            # same as build_noGC_images.sh
NFS="amf ausf udm udr pcf nrf nssf smf"          # same as build_noGC_images.sh
TAR=/local/aether/aether-$TAG.tar
VALUES=/local/aether/aether-sdcore-values.yaml

runtime=$(kubectl get nodes -o jsonpath='{.items[0].status.nodeInfo.containerRuntimeVersion}')
case "$runtime" in
  docker://*)     sudo docker load -i "$TAR" ;;
  containerd://*) sudo ctr -n k8s.io images import "$TAR" ;;
  *) echo "ERROR: unknown container runtime: $runtime"; exit 1 ;;
esac

SETS=""
for nf in $NFS; do
  SETS="$SETS --set 5g-control-plane.images.tags.$nf=5gc-$nf:$TAG"
done
helm -n aether-5gc upgrade sd-core oci://ghcr.io/omec-project/sd-core --version 6.1.2 -f "$VALUES" $SETS

for nf in $NFS; do
  kubectl -n aether-5gc rollout status deploy/$nf --timeout=300s
done
python3 /local/aether/aether_webui.py --namespace aether-5gc wait-nrf --timeout 300
kubectl -n aether-5gc get pods -o custom-columns=NAME:.metadata.name,IMAGE:.spec.containers[0].image
