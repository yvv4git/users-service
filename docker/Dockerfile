# multi-stage builds
FROM golang:1.14.0-alpine3.11 AS builder

RUN go version
RUN apk add git
RUN apk add build-base

COPY ./ /github.com/yvv4git/users-service
WORKDIR /github.com/yvv4git/users-service

RUN go mod download && go get -u ./...
RUN CGO_ENABLE=0 GOOS=linux go build -o ./bin/server.bin ./cmd/server/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root

COPY --from=0 /github.com/yvv4git/users-service/bin/server.bin .
COPY --from=0 /github.com/yvv4git/users-service/config ./config
COPY --from=0 /github.com/yvv4git/users-service/db ./db

EXPOSE 1234

CMD ["./server.bin"]
