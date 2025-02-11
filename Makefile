# Variables
PROJECT_NAME = authentication-svc
GATEWAY_NAME = gateway-svc
GO = go
PROTO_DIR = internal/proto
PROTO_FILE = $(PROTO_DIR)/auth.proto
WIN_PROTO_DIR = internal\proto
WIN_PROTO_FILE = $(WIN_PROTO_DIR)\auth.proto
PROTO_OUT = $(PROTO_DIR)
PORT = 50051

# Targets
.PHONY: build run proto clean

# Build the service
build:
	@echo "Building $(PROJECT_NAME)..."
	$(GO) build -o bin/$(PROJECT_NAME) ./cmd/main.go

# Run the service
run: build
	@echo "Running $(PROJECT_NAME) on port $(PORT)..."
	./bin/$(PROJECT_NAME)

# Generate Go code from .proto file
proto:
	@echo "Generating Go code from $(PROTO_FILE)..."
	protoc --go_out=$(PROTO_OUT) --go-grpc_out=$(PROTO_OUT) $(PROTO_FILE)
	cp $(PROTO_DIR)/auth.pb.go ../$(GATEWAY_NAME)/internal/proto/auth.pb.go
	cp $(PROTO_DIR)/auth_grpc.pb.go ../$(GATEWAY_NAME)/internal/proto/auth_grpc.pb.go

win-proto:
	@echo "Generating Go code from $(WIN_PROTO_FILE)..."
	protoc --go_out=$(WIN_PROTO_DIR) --go-grpc_out=$(WIN_PROTO_DIR) $(WIN_PROTO_FILE)
	xcopy $(WIN_PROTO_DIR)\auth.pb.go ..\$(GATEWAY_NAME)\internal\proto\auth.pb.go /i /Y
	xcopy $(WIN_PROTO_DIR)\auth_grpc.pb.go ..\$(GATEWAY_NAME)\internal\proto\auth_grpc.pb.go /i /Y

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf bin/$(PROJECT_NAME)