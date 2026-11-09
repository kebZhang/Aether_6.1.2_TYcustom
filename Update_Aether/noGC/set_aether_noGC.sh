#!/bin/bash
# Run on machine B (K8s node), AFTER apply_noGC_images.sh. Every helm upgrade
# puts amf's GOGC and smf's command back to the chart's, so run this again
# after every helm upgrade.
#   bash set_aether_noGC.sh aether-noGC_128_32.txt   # GOMAXPROCS from the preset + GC off
#   bash set_aether_noGC.sh gc-on                    # GC back on as deployed (GOMAXPROCS stays)
#   bash set_aether_noGC.sh                          # only print current values
#
# GC off = env vars on each NF Deployment:
#   GOGC=off           the runtime never collects on its own (no 2-minute periodic GC either)
#   GOMEMLIMIT=32GiB   backstop: the only point where it still collects on its own
#   GODEBUG=gctrace=1  one line per GC in kubectl logs; a forced GC ends with "(forced)"
# Two NFs need more, because the chart sets GOGC=200 itself:
#   amf  GOGC is in the chart's Deployment env (values config.gogc). set env
#        overwrites it; the next helm upgrade writes 200 back.
#   smf  smf-run.sh starts smf with "exec env GOTRACEBACK=crash GOGC=200 smf ...",
#        which overrides the pod's GOGC. The container command is patched to run a
#        copy of the script without "GOGC=200"; the next helm upgrade restores it.
set -euo pipefail

NS=aether-5gc
NFS="nrf amf ausf nssf pcf smf udm udr"
GC_OFF="GOGC=off GOMEMLIMIT=32GiB GODEBUG=gctrace=1"
SMF_CMD_GC_OFF='["sh","-c","sed s/GOGC=200// /opt/smf-run.sh > /tmp/smf-run.sh && exec sh /tmp/smf-run.sh"]'
SMF_CMD_CHART='["/opt/smf-run.sh"]'

plan=""                                          # lines "<nf> <kubectl set env args>"
if [ "${1:-}" = gc-on ]; then
  for nf in $NFS; do
    if [ "$nf" = amf ]; then plan+="amf GOGC=200 GOMEMLIMIT- GODEBUG-"$'\n'   # 200 = chart's config.gogc
    else                     plan+="$nf GOGC- GOMEMLIMIT- GODEBUG-"$'\n'; fi
  done
  smf_cmd=$SMF_CMD_CHART
elif [ $# -ge 1 ]; then
  while read -r nf val _; do
    case "$nf" in ''|\#*) continue ;; esac
    plan+="$nf GOMAXPROCS=$val $GC_OFF"$'\n'
  done < "$1"
  smf_cmd=$SMF_CMD_GC_OFF
fi

if [ -n "$plan" ]; then
  # NRF first, the other NFs after it
  for pass in nrf others; do
    changed=""
    while read -r nf args; do
      [ -z "$nf" ] && continue
      [ "$pass" = nrf ]    && [ "$nf" != nrf ] && continue
      [ "$pass" = others ] && [ "$nf" = nrf ]  && continue
      if [ "$nf" = smf ]; then
        kubectl -n $NS patch deploy/smf \
          -p "{\"spec\":{\"template\":{\"spec\":{\"containers\":[{\"name\":\"smf\",\"command\":$smf_cmd}]}}}}" </dev/null
      fi
      kubectl -n $NS set env deploy/$nf -c $nf $args </dev/null
      changed="$changed $nf"
    done <<< "$plan"
    for nf in $changed; do
      kubectl -n $NS rollout status deploy/$nf --timeout=300s
    done
  done
  python3 /local/aether/aether_webui.py --namespace $NS wait-nrf --timeout 300
fi

# What PID 1 (the NF binary: the run script exec's it) of every running pod has.
# For smf this is the GOGC it really runs with, after smf-run.sh's "env".
for nf in $NFS; do
  for pod in $(kubectl -n $NS get pods -l app=$nf -o go-template='{{range .items}}{{if not .metadata.deletionTimestamp}}{{.metadata.name}} {{end}}{{end}}'); do
    v=$(kubectl -n $NS exec $pod -c $nf -- sh -c 'tr "\0" "\n" </proc/1/environ | grep -E "^(GOMAXPROCS|GOGC|GOMEMLIMIT|GODEBUG)=" | tr "\n" " "')
    echo "$nf  $pod  ${v:-(none set)}"
  done
done
