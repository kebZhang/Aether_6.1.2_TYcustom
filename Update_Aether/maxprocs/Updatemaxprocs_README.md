# 设置每个 NF 的 GOMAXPROCS

xiaoxin 本机：

1.
```bash
cd /mnt/d/UB/Nemo/5GC_project/code/Aether_6.1.2_TYcustom/Update_Aether/maxprocs
```

2.
```bash
scp -i /home/tyz/.ssh/cloudlab ./set_aether_maxprocs.sh Tianyang@clnode123.clemson.cloudlab.us:/local/aether
```

3.
```bash
scp -i /home/tyz/.ssh/cloudlab ./aether-gomaxprocs*.txt Tianyang@clnode123.clemson.cloudlab.us:/local/aether
```

K8s 机器：

4. 进入 K8s 电脑

5.
```bash
cd /local/aether
```

6.
```bash
./set_aether_maxprocs.sh aether-gomaxprocs_128_32.txt
```

7. 检查每个NF的maxprocs
```bash
for nf in nrf amf ausf nssf pcf smf udm udr; do echo -n "$nf  "; kubectl -n aether-5gc exec deploy/$nf -c $nf -- sh -c 'tr "\0" "\n" </proc/1/environ | grep "^GOMAXPROCS=" || echo "GOMAXPROCS unset (Go default)"'; done
```
