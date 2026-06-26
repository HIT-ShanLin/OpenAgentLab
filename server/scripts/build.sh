#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SERVER_DIR="$(dirname "$SCRIPT_DIR")"
BUILD_DIR="${SERVER_DIR}/../out"

echo "==> Running tests..."
cd "$SERVER_DIR"
go test ./... -count=1

echo "==> Building all modules..."

MODULES=(
  "openagent-sandbox:./cmd/openagent-sandbox"
  "openagent-runtime:./cmd/openagent-runtime"
  "openagent-storage:./cmd/openagent-storage"
  "openagent-network:./cmd/openagent-network"
  "openagent-observability:./cmd/openagent-observability"
  "openagent-control-plane:./cmd/openagent-control-plane"
)

mkdir -p "$BUILD_DIR"

for entry in "${MODULES[@]}"; do
  name="${entry%%:*}"
  path="${entry##*:}"
  echo "  -> building ${name}..."
  CGO_ENABLED=0 go build -trimpath -o "${BUILD_DIR}/${name}" "$path"
done

echo "==> Build complete. Binaries:"
ls -lh "$BUILD_DIR"
