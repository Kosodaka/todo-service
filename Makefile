PHONY: generate
generate:
		mkdir internal\gRPC
		protoc --go_out ./internal/gRPC --go_opt=paths=source_relative \
			   --go-grpc_out ./internal/gRPC --go-grpc_opt=paths=source_relative \
					internal/proto/service_gRPC.proto