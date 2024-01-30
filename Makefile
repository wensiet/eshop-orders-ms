# Makefile

PROTO_ORDERS = proto/orders/orders.proto
PROTO_PRODUCTS = proto/products/products.proto

OUTPUT_DIR = ./gen/go/

PROTOC_ORDERS = protoc -I proto $(PROTO_ORDERS) \
    --go_out=$(OUTPUT_DIR) --go_opt=paths=source_relative \
    --go-grpc_out=$(OUTPUT_DIR) --go-grpc_opt=paths=source_relative

PROTOC_PRODUCTS = protoc -I proto $(PROTO_PRODUCTS) \
    --go_out=$(OUTPUT_DIR) --go_opt=paths=source_relative \
    --go-grpc_out=$(OUTPUT_DIR) --go-grpc_opt=paths=source_relative

orders:
	$(PROTOC_ORDERS)

products:
	$(PROTOC_PRODUCTS)