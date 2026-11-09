#!/bin/bash
# Run on machine B (K8s node) between runs: one full GC in every NF. Returns
# only after every GC has finished; exit 1 = do NOT start the next run.
#   bash force_gc_allNF.sh && <start the next run>
#
# GET :6060/gc (gc_trigger.go) inside each NF pod. Checks per NF:
#   - the request succeeded   otherwise: not the noGC image, or the pod is restarting
#   - gogc off                GC is off in the process (set_aether_noGC.sh; a helm
#                             upgrade turns it back on in amf and smf)
#   - auto_gc 0               no GC on its own since the pod started. >0 means
#                             GOMEMLIMIT was reached: every run since the last
#                             clean check is contaminated. Resets only on pod restart.
set -uo pipefail

NS=aether-5gc
NFS="nrf amf ausf nssf pcf smf udm udr"

val() { awk -v k="$1" '$1 == k { print $2 }' <<< "$r"; }
mib() { awk -v b="$1" 'BEGIN { printf "%.1f", b / 1048576 }'; }

FAIL=0
for nf in $NFS; do
  pods=$(kubectl -n $NS get pods -l app=$nf -o go-template='{{range .items}}{{if not .metadata.deletionTimestamp}}{{.metadata.name}} {{end}}{{end}}')
  [ -z "$pods" ] && { echo "[gc] $nf: FAILED -- no running pod" >&2; FAIL=1; continue; }
  for pod in $pods; do
    if ! r=$(kubectl -n $NS exec $pod -c $nf -- wget -q -O - -T 120 http://127.0.0.1:6060/gc); then
      echo "[gc] $nf: FAILED -- no reply from :6060/gc in $pod (not the noGC image, or restarting)" >&2
      FAIL=1; continue
    fi
    why=""
    [ "$(val gogc)" = off ]  || why+=" GC is ON (GOGC $(val gogc)): run set_aether_noGC.sh (again after every helm upgrade)."
    [ "$(val auto_gc)" = 0 ] || why+=" $(val auto_gc) automatic GC(s) since the pod started: runs since the last clean check are contaminated."
    printf "[gc] %-4s %-6s heap %s -> %s MiB\n" "$nf" "$([ -z "$why" ] && echo done || echo FAILED)" \
      "$(mib "$(val heap_before)")" "$(mib "$(val heap_after)")"
    [ -n "$why" ] && { echo "[gc] $nf:$why" >&2; FAIL=1; }
  done
done

if (( FAIL )); then
  echo "GC NOT CONFIRMED -- do NOT start the next run" >&2
  exit 1
fi
echo "ALL GC DONE ($NFS)"
