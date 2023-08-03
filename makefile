PROTO_DIR = examples/helloworld/helloworld
PROTO_SRC = $(PROTO_DIR)/helloworld.proto
PROTO_MESSAGES = $(PROTO_DIR)/helloworld.pb.go
PROTO_SERVICES = $(PROTO_DIR)/helloworld_grpc.pb.go
PROTO_DST = $(PROTO_MESSAGES) $(PROTO_SERVICES)

CLIENT_DIR = client
CLIENT_SRC = $(CLIENT_DIR)/main.go
CLIENT_DST = $(CLIENT_DIR)/client

SERVER_DIR = server
SERVER_SRC = $(SERVER_DIR)/main.go
SERVER_DST = $(SERVER_DIR)/server

GO_BUILD_OPTS = -tags netgo


.PHONY: all clean

all: $(CLIENT_DST) $(SERVER_DST)

clean:
	rm $(PROTO_DST) $(CLIENT_DST) $(SERVER_DST)

$(PROTO_MESSAGES): $(PROTO_SRC)
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		$^

$(PROTO_SERVICES): $(PROTO_SRC)
	protoc \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		$^

$(CLIENT_DST): $(CLIENT_SRC) $(PROTO_DST)
	cd $(CLIENT_DIR) && go build .

$(SERVER_DST): $(SERVER_SRC) $(PROTO_DST)
	cd $(SERVER_DIR) && go build .
