package main

import (
	"fmt"
	"log"
	"net"

	"github.com/yvv4git/users-service/internal/api"
	"github.com/yvv4git/users-service/internal/servers"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Server")

	// Create and register grpc server.
	grpcServer := grpc.NewServer()
	usersServer := servers.NewUsersServer()
	api.RegisterUsersServer(grpcServer, usersServer)

	// Bind port.
	tcpListner, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	// Start server.
	if err = grpcServer.Serve(tcpListner); err != nil {
		log.Fatal(err)
	}
}
