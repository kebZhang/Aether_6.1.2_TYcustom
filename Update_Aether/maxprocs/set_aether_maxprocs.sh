#!/bin/bash
# Run on machine B (K8s node). Sets GOMAXPROCS on each Aether NF Deployment from a
# preset file, waits for the pods to restart, then prints the GOMAXPROCS each
# running NF process actually has.
#   bash set_aether_maxprocs.sh aether-gomaxprocs_128_32.txt   # apply a preset
#   bash set_aether_maxprocs.sh                                # only print current values
#
# Preset line: "<nf> <value>"; value "-" removes GOMAXPROCS (back to Go's default).
# kubectl set env only restarts a Deployment whose value actually changes.
set -euo pipefail

NS=aether-5gc
NFS="nrf amf ausf nssf pcf smf udm udr"

if [ $# -ge 1 ]; then
  # NRF first, the other NFs after it
  for pass in nrf others; do
    changed=""
    while read -r nf val _; do
      case "$nf" in ''|\#*) continue ;; esac
      [ "$pass" = nrf ]    && [ "$nf" != nrf ] && continue
      [ "$pass" = others ] && [ "$nf" = nrf ]  && continue
      if [ "$val" = "-" ]; then arg="GOMAXPROCS-"; else arg="GOMAXPROCS=$val"; fi
      kubectl -n $NS set env deploy/$nf -c $nf "$arg" </dev/null
      changed="$changed $nf"
    done < "$1"
    for nf in $changed; do
      kubectl -n $NS rollout status deploy/$nf --timeout=300s
    done
  done
  python3 /local/aether/aether_webui.py --namespace $NS wait-nrf --timeout 300
fi

# Read GOMAXPROCS from PID 1 (the NF binary: the run script exec's it) of every
# pod that is not terminating.
for nf in $NFS; do
  for pod in $(kubectl -n $NS get pods -l app=$nf -o go-template='{{range .items}}{{if not .metadata.deletionTimestamp}}{{.metadata.name}} {{end}}{{end}}'); do
    v=$(kubectl -n $NS exec $pod -c $nf -- sh -c 'tr "\0" "\n" </proc/1/environ | grep "^GOMAXPROCS=" || echo "GOMAXPROCS unset (Go default)"')
    echo "$nf  $pod  $v"
  done
done
