SERVER := ./server/main.go
CLIENT := ./client/main.go

CLIENT_BINARY := uno_client
SERVER_BINARY := uno_server

PKG := $(shell find pkg -type f)

PROTO_DIR := ./proto
PROTO_OUT := ./pkg/proto

all: protos client server

protos: $(patsubst $(PROTO_DIR)/%.proto,$(PROTO_OUT)/%.pb.go, $(wildcard $(PROTO_DIR)/*.proto))
client: protos $(CLIENT_BINARY)
server: protos $(SERVER_BINARY)

$(CLIENT_BINARY): $(CLIENT) $(PKG)
	go build -o $(CLIENT_BINARY) $(CLIENT)

$(SERVER_BINARY): $(SERVER) $(PKG)
	go build -o $(SERVER_BINARY) $(SERVER)

$(PROTO_OUT)/%.pb.go: $(wildcard $(PROTO_DIR)/*.proto)
	protoc --go_out=proto --go-grpc_out=proto proto/uno.proto

clean:
	go clean
	rm -f $(CLIENT_BINARY) $(SERVER_BINARY) $(wildcard $(PROTO_OUT)/*.pb.go)
