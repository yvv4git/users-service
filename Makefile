gen_protoc_service_go:
	protoc -I protocol/ --go_out=plugins=grpc:service/internal/rpc/ protocol/users.proto
