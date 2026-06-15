# Ckyn - Lightweight Container Security Assessment Tool

A lightweight container security assessment tool for penetration testing and security research.

## Features

- Container security evaluation
- Exploit execution
- Built-in security tools (netstat, ps, vi, etc.)
- K8s/Docker API interaction

## Usage

```bash
# Evaluate container security
./ckyn evaluate

# List available exploits
./ckyn run --list

# Run specific exploit
./ckyn run <exploit-name>

# Use built-in tools
./ckyn netstat
./ckyn ps
./ckyn vi <file>
```

## Build

```bash
GOOS=linux GOARCH=amd64 go build ./cmd/ckyn/
```

## Legal Disclaimer

This tool is for security testing purposes only. Usage against targets without prior consent is illegal.
