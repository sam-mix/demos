PROTOC=protoc
PROTO_FILES=$(wildcard *.proto)
GO_OUT_DIR=.
GRPC_OUT_DIR=.

.PHONY: all go grpc clean

all: go grpc

go: $(PROTO_FILES)
    $(PROTOC) --go_out=$(GO_OUT_DIR) $(PROTO_FILES)

grpc: $(PROTO_FILES)
    $(PROTOC) --go-grpc_out=$(GRPC_OUT_DIR) $(PROTO_FILES)

clean:
    rm -rf $(GO_OUT_DIR) $(GRPC_OUT_DIR)