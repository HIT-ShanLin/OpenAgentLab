.PHONY: all build clean test lint sandbox runtime storage network observability control-plane

SERVER_DIR = server
BUILD_DIR = out

# Build all modules
all: sandbox runtime storage network observability control-plane

sandbox:
	cd $(SERVER_DIR) && CGO_ENABLED=0 go build -trimpath -o ../$(BUILD_DIR)/openagent-sandbox ./openagent-sandbox/cmd

runtime:
	cd $(SERVER_DIR) && CGO_ENABLED=0 go build -trimpath -o ../$(BUILD_DIR)/openagent-runtime ./openagent-runtime/cmd

storage:
	cd $(SERVER_DIR) && CGO_ENABLED=0 go build -trimpath -o ../$(BUILD_DIR)/openagent-storage ./openagent-storage/cmd

network:
	cd $(SERVER_DIR) && CGO_ENABLED=0 go build -trimpath -o ../$(BUILD_DIR)/openagent-network ./openagent-network/cmd

observability:
	cd $(SERVER_DIR) && CGO_ENABLED=0 go build -trimpath -o ../$(BUILD_DIR)/openagent-observability ./openagent-observability/cmd

control-plane:
	cd $(SERVER_DIR) && CGO_ENABLED=0 go build -trimpath -o ../$(BUILD_DIR)/openagent-control-plane ./openagent-control-plane/cmd

# Run all tests
test:
	cd $(SERVER_DIR) && go test ./... -count=1 -race

# Run lint (requires golangci-lint)
lint:
	cd $(SERVER_DIR) && golangci-lint run ./...

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR)
