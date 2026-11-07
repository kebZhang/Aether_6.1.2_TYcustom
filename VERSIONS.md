# Aether SD-Core 官方源码（release 6.1.2）

下载日期：2026-11-04。来源：GitHub `omec-project/*` 各仓库的 release tag 源码包（不含 `.git`）。

版本依据：Aether OnRamp `vars/main.yml` 固定的 `oci://ghcr.io/omec-project/sd-core` **6.1.2**，
即 `cloudlab/K8s/Deploy_Aether/Deploy_Aether_K8s.sh` 部署的那个 chart。
每个 NF 的 tag 取自 `sdcore-helm-charts/5g-control-plane/values.yaml` 的 `images.tags`
（镜像 `rel-X.Y.Z` ↔ 仓库 tag `vX.Y.Z`），各目录里的 `VERSION` 文件已核对一致。

## NFs/

| 目录 | 仓库 | tag | 集群里的镜像 |
|---|---|---|---|
| af | omec-project/af | v1.0.0 | 5gc-af:rel-1.0.0（Deploy_Aether 里关掉了） |
| amf | omec-project/amf | v3.3.2 | 5gc-amf:rel-3.3.2 |
| ausf | omec-project/ausf | v3.1.2 | 5gc-ausf:rel-3.1.2 |
| nrf | omec-project/nrf | v3.1.2 | 5gc-nrf:rel-3.1.2 |
| nssf | omec-project/nssf | v3.1.2 | 5gc-nssf:rel-3.1.2 |
| pcf | omec-project/pcf | v3.2.2 | 5gc-pcf:rel-3.2.2 |
| smf | omec-project/smf | v4.3.3 | 5gc-smf:rel-4.3.3 |
| udm | omec-project/udm | v3.1.2 | 5gc-udm:rel-3.1.2 |
| udr | omec-project/udr | v3.1.2 | 5gc-udr:rel-3.1.2 |
| webui | omec-project/webconsole | v3.1.4 | 5gc-webui:rel-3.1.4 |
| metricfunc | omec-project/metricfunc | v2.1.3 | metricfunc:rel-2.1.3（未部署） |
| sctplb | omec-project/sctplb | v1.9.4 | sctplb:rel-1.9.4（未部署） |
| upfadapter | omec-project/upfadapter | v2.2.1 | upfadapter:rel-2.2.1（未部署） |
| pod-init | omec-project/aether-pod-init | v1.3.1 | pod-init:rel-1.3.1（init 容器） |

## shared-libs/

各 NF `go.mod` 里引用的 omec-project 库，每个用到的版本各一份（NF 的 go.mod 都没有 replace）：

| 目录 | 被谁引用 |
|---|---|
| openapi@v2.2.4 | amf ausf nrf nssf pcf smf udm udr |
| openapi@v2.2.5 | webui metricfunc |
| openapi@v2.4.1 | af |
| util@v1.8.10 | amf ausf nrf nssf pcf smf udm udr |
| util@v1.8.13 | webui metricfunc |
| util@v1.9.3 | af |
| nas@v2.2.5 | amf smf |
| ngap@v2.1.7 | amf smf |
| ngap@v2.1.9 | sctplb |

注意：这里的 shared-libs 只是源码副本。NF 编译时默认按 go.mod 从 Go proxy 拉库，
改了这里的库要生效，必须在对应 NF 的 go.mod 里加 `replace => <本地路径>`。

## sdcore-helm-charts/

omec-project/sdcore-helm-charts v6.1.2（chart 源码，与 ghcr.io 上的 sd-core 6.1.2 同一版本）。

## 没有下载的

simapp（omec-sub-provision）、bess-upf、gnbsim：属于同一个 release，但 Deploy_Aether 的部署里都关掉了。
