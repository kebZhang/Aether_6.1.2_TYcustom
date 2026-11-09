#!/bin/bash
# Run on machine B (K8s node). Imports the tar from build_aether_images.sh and
# points the sd-core release at the new tags.
#   ./apply_aether_images.sh <tar>        e.g. ./apply_aether_images.sh aether-NF8HTTP_1108.tar
set -euo pipefail

TAR=${1:?usage: ./apply_aether_images.sh <aether-TAG.tar>}
NFS="amf ausf udm udr pcf nrf nssf smf"           # same as build_aether_images.sh
VALUES=/local/aether/aether-sdcore-values.yaml   # uploaded by Deploy_Aether_K8s.sh

# TAG comes from the file name build_aether_images.sh gave the tar: aether-<TAG>.tar
name=$(basename "$TAR")
case "$name" in
  aether-?*.tar) TAG=${name#aether-}; TAG=${TAG%.tar} ;;
  *) echo "ERROR: expected aether-<TAG>.tar (as written by build_aether_images.sh), got: $name"; exit 1 ;;
esac

# Every NF in NFS must be in the tar. A tar built when NFS had only 6 NFs
# (e.g. aether-NF8HTTP_1108.tar) has no nssf/smf: stop here, before helm upgrade.
manifest=$(tar -xOf "$TAR" manifest.json)
for nf in $NFS; do
  grep -q "5gc-$nf:$TAG\"" <<< "$manifest" || { echo "ERROR: $name has no 5gc-$nf:$TAG -- rebuild it with build_aether_images.sh"; exit 1; }
done

# 1. Import into the node's container runtime
runtime=$(kubectl get nodes -o jsonpath='{.items[0].status.nodeInfo.containerRuntimeVersion}')
case "$runtime" in
  docker://*)     sudo docker load -i "$TAR" ;;
  containerd://*) sudo ctr -n k8s.io images import "$TAR" ;;
  *) echo "ERROR: unknown container runtime: $runtime"; exit 1 ;;
esac

# 2. helm upgrade: same chart / version / values as Deploy_Aether_K8s.sh, plus the new tags.
#    pullPolicy stays IfNotPresent, so the imported images are used without pulling.
SETS=""
for nf in $NFS; do
  SETS="$SETS --set 5g-control-plane.images.tags.$nf=5gc-$nf:$TAG"
done
helm -n aether-5gc upgrade sd-core oci://ghcr.io/omec-project/sd-core --version 6.1.2 -f "$VALUES" $SETS

# 3. Wait for the new pods, then for NRF registration
for nf in $NFS; do
  kubectl -n aether-5gc rollout status deploy/$nf --timeout=300s
done
python3 /local/aether/aether_webui.py --namespace aether-5gc wait-nrf --timeout 300
kubectl -n aether-5gc get pods -o custom-columns=NAME:.metadata.name,IMAGE:.spec.containers[0].image
