#!/bin/bash
# Run on machine A (build machine). Builds one image per NF in NFS from the
# Aether_6.1.2_TYcustom checkout this script lives in, and saves them all into one tar.
set -euo pipefail

TAG=NF8HTTP_1108                                 # new tag per build; apply_aether_images.sh must use the same
NFS="amf ausf udm udr pcf nrf"                   # keep only the NFs you changed
OUT=/local/aether-$TAG.tar                       # outside the repo, so git never sees it

# Images are named ghcr.io/omec-project/5gc-<nf>:<TAG> because the chart builds
# every image as <images.repository><images.tags.X> with one shared repository
# prefix (the init container uses it too), so only the tag is overridden on B.
HERE="$(cd "$(dirname "$0")" && pwd)"            # <repo>/Update_Aether
SRC="$(dirname "$HERE")"                         # repo root = build context

# Install Docker (official get.docker.com script, includes BuildKit) if missing
if ! command -v docker >/dev/null; then
  curl -fsSL https://get.docker.com | sudo sh
fi
# Use sudo when the current user is not in the docker group
DOCKER=docker
docker info >/dev/null 2>&1 || DOCKER="sudo DOCKER_BUILDKIT=1 docker"

export DOCKER_BUILDKIT=1
IMAGES=""
for nf in $NFS; do
  bin=$nf
  [ "$nf" = webui ] && bin=webconsole            # chart runs "webconsole -cfg ..."
  img=ghcr.io/omec-project/5gc-$nf:$TAG
  $DOCKER build -f "$HERE/Dockerfile.nf" --build-arg NF=$nf --build-arg BIN=$bin -t "$img" "$SRC"
  IMAGES="$IMAGES $img"
done

$DOCKER save $IMAGES -o "$OUT"
sudo chmod 644 "$OUT"                            # readable for scp when saved via sudo
ls -lh "$OUT"
