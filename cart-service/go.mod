module github.com/wazwki/WearStore/cart-service

go 1.23.2

require (
	github.com/joho/godotenv v1.5.1
	github.com/redis/go-redis/v9 v9.7.0
	github.com/wazwki/WearStore/product-service v0.0.0-20241214040947-96234c8a1fe0
	github.com/wazwki/WearStore/user-service v0.0.0-20241215201653-ebc3c20eac8c
	go.uber.org/zap v1.27.0
	google.golang.org/grpc v1.69.0
	google.golang.org/protobuf v1.35.2
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
)
