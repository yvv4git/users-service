gen_protoc_service_go:
	protoc -I protocol/ --go_out=plugins=grpc:internal/api/ protocol/users.proto
