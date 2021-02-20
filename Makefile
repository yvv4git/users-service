gen_protoc_service_go:
	protoc -I protocol/ --go_out=plugins=grpc:internal/api/ protocol/users.proto

gen_mocks:
	go generate -v ./...

tests:
	go test -v ./...

build:
	GOOS=linux GOARCH=amd64 go build -o server.bin cmd/server/main.go

migrations_create:
	sql-migrate new -env="development" create_tb_users

migrations_up:
	sql-migrate up -env="development"

migrations_down:
	sql-migrate down -env="development" -dryrun

docker_build:
	docker build --no-cache --rm -m 2G -t yvv4docker/users-service -f docker/Dockerfile .

docker_run:
	docker run --rm --name users-service -p 1234:1234 yvv4docker/users-service