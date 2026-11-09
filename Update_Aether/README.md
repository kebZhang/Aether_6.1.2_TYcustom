# 改了 Aether 代码之后：编译镜像 → 部署到 K8s

- 机器 A：编译机，需要能上网（没装 docker 的话，build 脚本会自动安装）
- 机器 C：xiaoxin，本地 WSL，负责在 A 和 B 之间中转文件
- 机器 B：K8s 节点，已经用 `Deploy_Aether_K8s.sh` 部署过 Aether，所以 `/local/aether/` 下已有 `aether-sdcore-values.yaml` 和 `aether_webui.py`

本文件夹在仓库里，A 上 `git pull` 时会随代码一起拿到，不用另外传。

| 文件 | 在哪台机器用 | 作用 |
|---|---|---|
| `Dockerfile.nf` | A | 编译单个 NF 的镜像。build context 是仓库根目录 |
| `build_aether_images.sh` | A | 参数是 tag。逐个编译 `NFS` 里列的 NF，最后打成一个 tar 包 `/local/aether-<tag>.tar` |
| `apply_aether_images.sh` | B | 参数是 tar 包。把镜像导入 K8s，执行 helm upgrade 换成新 tag，再等 NF 重新注册到 NRF |

下面以 tag `NF8HTTP_1108` 为例，`<A>`、`<B>` 换成两台机器的地址。

## 1. C：改好 NFS，提交代码

打开 `build_aether_images.sh` 和 `apply_aether_images.sh`，把开头的 `NFS` 改成一样的值：只填你改过的 NF。

tag 不用改脚本，第 2 步编译时作为参数传入。

```bash
cd /mnt/d/UB/Nemo/5GC_project/code/Aether_6.1.2_TYcustom
git add -A && git commit -m "..." && git push
```

## 2. A：拉代码、编译

```bash
cd /local
git clone https://github.com/kebZhang/Aether_6.1.2_TYcustom.git   # 第一次
# 以后： cd /local/Aether_6.1.2_TYcustom && git pull
cd /local/Aether_6.1.2_TYcustom/Update_Aether
./build_aether_images.sh NF8HTTP_1108
```

参数就是 tag，每次编译都用一个新的 tag。编译完会生成 `/local/aether-NF8HTTP_1108.tar`。

## 3. C：从 A 下载 tar，再上传到 B

```bash
cd /mnt/d/UB/Nemo/5GC_project/code/Aether_6.1.2_TYcustom/Update_Aether
scp -i ~/.ssh/cloudlab Tianyang@<A>:/local/aether-NF8HTTP_1108.tar /tmp/
scp -i ~/.ssh/cloudlab /tmp/aether-NF8HTTP_1108.tar apply_aether_images.sh Tianyang@<B>:/local/aether/
```

不要给 tar 改名：B 上的脚本从文件名 `aether-NF8HTTP_1108.tar` 里取出 tag `NF8HTTP_1108`。

## 4. B：导入镜像并更新

```bash
cd /local/aether
./apply_aether_images.sh aether-NF8HTTP_1108.tar
```

最后一行会打印每个 pod 用的镜像，确认改过的 NF 显示的是 `...:NF8HTTP_1108`。之后照常进 PacketRusher pod 运行 `./entry_aether.sh`。

## 注意

- **改了 `shared-libs/` 里的库**：要在每个用到它的 NF 的 `go.mod` 里加一行 `replace`，例如 `replace github.com/omec-project/openapi/v2 => ../../shared-libs/openapi@v2.2.4`，否则编译时用的仍是从网上下载的官方版本，改动不会生效。只改 `NFs/` 下的代码不用加。
- **回退到官方镜像**：在 B 上执行一次 helm upgrade，不带 `--set`：
  `helm -n aether-5gc upgrade sd-core oci://ghcr.io/omec-project/sd-core --version 6.1.2 -f /local/aether/aether-sdcore-values.yaml`
  重新跑一次 `Deploy_Aether_K8s.sh` 也会回到官方镜像。
- **多节点集群**：脚本只把镜像导入了 B 这一台。如果 NF 的 pod 可能调度到别的节点，每个节点都要导入一次。
