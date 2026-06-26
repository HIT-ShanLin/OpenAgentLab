.PHONY: all build clean test lint sandbox runtime storage network observability control-plane

SERVER_DIR = server
BUILD_DIR = out

# Build all modules
all: sandbox runtime storage network observability control-plane

sandbox:
	cd $(SERVER_DIR) && CGO_ENABLED=0 go build -trimpath -o ../$(BUILD_DIR)/openagent-sandbox ./cmd/openagent-sandbox

runtime:
	cd $(SERVER_DIR) && CGO_ENABLED=0 go build -trimpath -o ../$(BUILD_DIR)/openagent-runtime ./cmd/openagent-runtime

storage:
	cd $(SERVER_DIR) && CGO_ENABLED=0 go build -trimpath -o ../$(BUILD_DIR)/openagent-storage ./cmd/openagent-storage

network:
	cd $(SERVER_DIR) && CGO_ENABLED=0 go build -trimpath -o ../$(BUILD_DIR)/openagent-network ./cmd/openagent-network

observability:
	cd $(SERVER_DIR) && CGO_ENABLED=0 go build -trimpath -o ../$(BUILD_DIR)/openagent-observability ./cmd/openagent-observability

control-plane:
	cd $(SERVER_DIR) && CGO_ENABLED=0 go build -trimpath -o ../$(BUILD_DIR)/openagent-control-plane ./cmd/openagent-control-plane

# Run all tests
test:
	cd $(SERVER_DIR) && go test ./... -count=1 -race

# Run lint (requires golangci-lint)
lint:
	cd $(SERVER_DIR) && golangci-lint run ./...

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR)
