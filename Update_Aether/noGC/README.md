# noGC：128_32 配置下，实验过程中所有 NF 都不做 GC

分两部分：

- **关 GC：只靠设置。** 给每个 NF 的 pod 设 `GOGC=off`，进程从启动起就不会自己做 GC。
- **实验间隙强制 GC：需要改代码。** 镜像在 NF8HTTP 的代码上只多了一个文件 `NFs/<nf>/gc_trigger.go`（8 个 NF 内容相同）。它只在 `:6060/gc` 上等待请求，收到请求就做一次完整的 GC，做完才回复。它不改任何 GC 设置：不设 `GOGC=off` 时，这个镜像的 GC 和 NF8HTTP 镜像完全一样。

镜像里有 8 个 NF：amf ausf udm udr pcf nrf 用 NF8HTTP 的代码；nssf 和 smf 原来用官方镜像，这次用官方源码加 `gc_trigger.go` 编译（官方 Dockerfile 用的也是 golang:1.27.1 + alpine:3.24，`CGO_ENABLED=0`，和 `../Dockerfile.nf` 一样），没有 8 连接的改动。

## 文件

| 文件 | 在哪台机器用 | 作用 |
|---|---|---|
| `build_noGC_images.sh` | A | 编译 8 个 NF 的镜像，打成 `/local/aether-NF8HTTP_noGC_1109.tar` |
| `apply_noGC_images.sh` | B | 导入镜像，helm upgrade 换成这个 tag。只换镜像，不管 GC |
| `set_aether_noGC.sh` | B | 按预设设置每个 NF 的 GOMAXPROCS，并关掉 GC |
| `aether-noGC_128_32.txt` | B | 预设：amf 128，其余 NF 32 |
| `force_gc_allNF.sh` | B | 每个 NF 做一次 GC，并检查 GC 确实关着、没有自动 GC 过 |

## `set_aether_noGC.sh` 怎么关 GC

给每个 NF 的 Deployment 设三个环境变量：

| 变量 | 作用 |
|---|---|
| `GOGC=off` | 进程不会自己做 GC（Go 每 2 分钟一次的定时 GC 也一起关掉） |
| `GOMEMLIMIT=32GiB` | 保底。GC 关掉后，只有进程内存涨到这个值时 Go 才会自己做 GC，防止把节点内存用光。单位必须写 `GiB`，写 `32GB` 进程会启动失败 |
| `GODEBUG=gctrace=1` | 每次 GC 在 `kubectl logs` 里打一行，强制 GC 的行末尾是 `(forced)` |

有两个 NF 的 `GOGC=200` 是 chart 自己写的，要另外处理：

- **amf**：chart 在 Deployment 的环境变量里写了 `GOGC`（来自 values 的 `config.gogc`，默认 `"200"`）。脚本用 `kubectl set env` 把它改成 `off`。
- **smf**：启动脚本是 `exec env GOTRACEBACK=crash GOGC=200 smf ...`，pod 上设的 `GOGC` 会被这里覆盖。脚本用 `kubectl patch` 把 smf 容器的启动命令改成：先把启动脚本复制到 `/tmp` 并删掉其中的 `GOGC=200`，再执行这份副本。这样 smf 也和其他 NF 一样，用 pod 上的 `GOGC`。

**每次 helm upgrade 都会把 amf 的 `GOGC` 改回 200、把 smf 的启动命令改回原样**（这两个都是 chart 管的）。其余的变量 chart 不管，helm upgrade 后仍然保留。所以：先跑 `apply_noGC_images.sh`，再跑 `set_aether_noGC.sh`；以后每次 helm upgrade 之后都要重新跑一次 `set_aether_noGC.sh`。忘了也没关系，`force_gc_allNF.sh` 会报 `GC is ON`。

## 使用步骤

### 1. C：提交代码

```bash
cd /mnt/d/UB/Nemo/5GC_project/code/Aether_6.1.2_TYcustom
git add -A && git commit -m "noGC: gc_trigger.go -- 1109" && git push
```

### 2. A：编译

```bash
cd /local/Aether_6.1.2_TYcustom && git pull
bash Update_Aether/noGC/build_noGC_images.sh
```

### 3. C：从 A 下载 tar，连同脚本传到 B

```bash
cd /mnt/d/UB/Nemo/5GC_project/code/Aether_6.1.2_TYcustom/Update_Aether/noGC
scp -i ~/.ssh/cloudlab Tianyang@<A>:/local/aether-NF8HTTP_noGC_1109.tar /tmp/
scp -i ~/.ssh/cloudlab /tmp/aether-NF8HTTP_noGC_1109.tar *.sh aether-noGC_128_32.txt Tianyang@<B>:/local/aether/
```

### 4. B：换镜像、关 GC、确认

```bash
cd /local/aether
bash apply_noGC_images.sh
bash set_aether_noGC.sh aether-noGC_128_32.txt
bash force_gc_allNF.sh
```

`set_aether_noGC.sh` 最后打印每个 NF 进程实际拿到的值，8 个 NF 都应该是 `GOGC=off`，例如：

```
amf  amf-7c4b...  GOMAXPROCS=128 GOGC=off GOMEMLIMIT=32GiB GODEBUG=gctrace=1
smf  smf-5d8f...  GOMAXPROCS=32 GOGC=off GOMEMLIMIT=32GiB GODEBUG=gctrace=1
...
```

smf 显示 `GOGC=200` 说明启动命令没有改成功。

`force_gc_allNF.sh` 最后一行是 `ALL GC DONE`，就说明 8 个 NF 都是 noGC 镜像，并且 GC 都关着。

### 5. 实验：每个 run 之前在 B 上跑一次

```bash
bash /local/aether/force_gc_allNF.sh
```

看到 `ALL GC DONE` 再进 PacketRusher pod 运行 `./entry_aether.sh`。最后一个 run 结束后也跑一次，用来确认最后一个 run 里没有发生 GC。

每个 NF 输出一行，例如：

```
[gc] udr  done   heap 412.3 -> 18.0 MiB
```

`heap` 前一个数是上一个 run 留在堆上的内存（GC 关着，run 里分配的都还在），后一个数是 GC 后剩下的。

出现 `FAILED` 时，最后一行是 `GC NOT CONFIRMED -- do NOT start the next run`，退出码为 1，不要开始下一个 run。原因写在 `FAILED` 下面一行：

| 提示 | 原因 |
|---|---|
| `no reply from :6060/gc` | 这个 NF 不是 noGC 镜像，或者 pod 正在重启 |
| `GC is ON (GOGC 200)` | 没跑 `set_aether_noGC.sh`，或者之后又做过 helm upgrade（amf、smf 会被改回 200） |
| `N automatic GC(s) since the pod started` | 这个 NF 自己做过 GC，说明内存涨到了 GOMEMLIMIT。上一次检查通过之后的 run 都受影响。这个计数只有 pod 重启才清零 |

也可以直接看 gctrace：`kubectl -n aether-5gc logs deploy/udr | grep '^gc '`，每一行末尾都应该是 `(forced)`。

## 恢复 GC

```bash
bash set_aether_noGC.sh gc-on
```

删掉 `GOGC`、`GOMEMLIMIT`、`GODEBUG`（amf 的 `GOGC` 改回 chart 的 200），smf 的启动命令改回原样，GOMAXPROCS 保持不变。镜像不用换：noGC 镜像在 GC 开着时和 NF8HTTP 镜像一样，只是多一个没人访问的 `:6060`。

要换回 NF8HTTP 镜像：运行 `../apply_aether_images.sh`。它只设 6 个 NF 的 tag，nssf 和 smf 会回到 chart 默认的官方镜像。

## 注意

- **换回 NF8HTTP 镜像后，没跑 `gc-on` 的话 GC 仍然是关的**：除了 amf、smf 被 helm upgrade 改回 200，其余 6 个 NF 的 `GOGC=off` 还在，而旧镜像没有 `:6060/gc`，没法强制 GC。不确定当前状态时，不带参数运行 `bash set_aether_noGC.sh` 查看。
- **GOMEMLIMIT 是每个进程单独算的**：8 × 32GiB = 256GiB。如果节点内存比这小，多个 NF 同时涨到很大时节点会先没内存。正常情况下一个 run 的分配量远小于 32GiB，可以从 `force_gc_allNF.sh` 输出的 `heap` 前一个数看出每个 run 实际分配了多少。
- **强制 GC 只回收，不把内存还给操作系统**，回收的内存留在进程里给下一个 run 用，所以下一个 run 不会因为重新向系统要内存而变慢。
- `:6060` 只有 `/gc` 一个路径，不是 pprof。amf 和 smf 本来就在 `:5001` 上开着 pprof（chart 的 `debugProfilePort`），和这里无关。
- webui 不是 NF，不在列表里，它的 GC 照常运行。
