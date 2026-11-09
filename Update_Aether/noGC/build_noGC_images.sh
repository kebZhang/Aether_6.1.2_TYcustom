#!/bin/bash
# Run on machine A (build machine). Same as ../build_aether_images.sh, for the
# noGC image: the NF8HTTP code plus NFs/<nf>/gc_trigger.go (forced GC on
# :6060/gc; GC itself stays as the pod env sets it) in every NF below.
# nssf and smf have no NF8HTTP change; they are built only to get gc_trigger.go.
set -euo pipefail

TAG=NF8HTTP_noGC_1109                            # apply_noGC_images.sh must use the same
NFS="amf ausf udm udr pcf nrf nssf smf"          # every NF that must be force-GC'd
OUT=/local/aether-$TAG.tar

HERE="$(cd "$(dirname "$0")/.." && pwd)"         # <repo>/Update_Aether (Dockerfile.nf)
SRC="$(dirname "$HERE")"                         # repo root = build context

if ! command -v docker >/dev/null; then
  curl -fsSL https://get.docker.com | sudo sh
fi
DOCKER=docker
docker info >/dev/null 2>&1 || DOCKER="sudo DOCKER_BUILDKIT=1 docker"

export DOCKER_BUILDKIT=1
IMAGES=""
for nf in $NFS; do
  [ -f "$SRC/NFs/$nf/gc_trigger.go" ] || { echo "ERROR: NFs/$nf/gc_trigger.go missing (git pull?)"; exit 1; }
  img=ghcr.io/omec-project/5gc-$nf:$TAG
  $DOCKER build -f "$HERE/Dockerfile.nf" --build-arg NF=$nf --build-arg BIN=$nf -t "$img" "$SRC"
  IMAGES="$IMAGES $img"
done

$DOCKER save $IMAGES -o "$OUT"
sudo chmod 644 "$OUT"
ls -lh "$OUT"
