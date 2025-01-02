module github.com/wazwki/WearStore/cart-service

go 1.23.2

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.25.1
	github.com/redis/go-redis/v9 v9.7.0
	github.com/wazwki/WearStore/product-service v0.0.0-20241214040947-96234c8a1fe0
	github.com/wazwki/WearStore/user-service v0.0.0-20241215201653-ebc3c20eac8c
	go.uber.org/zap v1.27.0
	google.golang.org/genproto/googleapis/api v0.0.0-20241219192143-6b3ec007d9bb
	google.golang.org/grpc v1.69.2
	google.golang.org/protobuf v1.36.1
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241219192143-6b3ec007d9bb // indirect
)
