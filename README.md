# OpenAgentLab

> Agent Sandbox Infrastructure — building the operating system for AI agents.

OpenAgentLab is a monorepo providing foundational infrastructure for agent execution environments: sandbox isolation, virtual networking, ephemeral storage, runtime engines, observability, and control-plane orchestration.

## Architecture

```
Agent Request → Gateway → Scheduler → Sandbox → Execution → Storage → Result
```

## Monorepo Layout

```
OpenAgentLab/
├── server/                        # Go backend (DDD / clean architecture)
│   ├── cmd/                       # Entry points (one per module)
│   │   ├── openagent-sandbox/
│   │   ├── openagent-runtime/
│   │   ├── openagent-storage/
│   │   ├── openagent-network/
│   │   ├── openagent-observability/
│   │   └── openagent-control-plane/
│   ├── internal/                  # Private packages, layered by module
│   │   ├── sandbox/               #   domain → service → repo → adapter
│   │   ├── runtime/
│   │   ├── storage/
│   │   ├── network/
│   │   ├── observability/
│   │   └── controlplane/
│   └── pkg/                       # Shared utilities (logger, config)
├── api/proto/                     # Protobuf definitions (shared)
├── web/                           # Frontend (reserved)
└── out/                           # Build artifacts
```

## Modules

| Module | Description | Port |
|--------|-------------|------|
| `openagent-sandbox` | Linux Namespace / cgroup sandbox isolation | 9100 |
| `openagent-runtime` | Container & MicroVM execution environments | 9101 |
| `openagent-storage` | Ephemeral block devices, OverlayFS, shared volumes | 9102 |
| `openagent-network` | Virtual networking (veth, bridge, iptables) | 9103 |
| `openagent-observability` | Prometheus metrics, structured logging | 9104 |
| `openagent-control-plane` | Task scheduling, queuing, gateway routing | 9105 |

## Quick Start

```bash
# Build all modules
make all

# Run tests
make test

# Build a single module
make sandbox

# Clean
make clean
```

## Design Principles

- **DDD layered architecture**: `domain → service → repo interface → adapter`
- **Manual dependency injection**: explicit wiring in `bootstrap/`, no framework magic
- **Interface-driven adapters**: kernel primitives (namespace, cgroup, bridge) behind contracts
- **Shared nothing between modules**: each module owns its domain and can be deployed independently

## Roadmap

See [[01_项目发展路径探讨]] (Obsidian) for the full 12-month plan — from mini-sandbox to a complete Agent Infra platform.

## License

MIT
