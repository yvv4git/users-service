package main

import (
	"fmt"
	"log"

	/* "github.com/yvv4git/users-service/internal/api" */
	"github.com/yvv4git/users-service/internal/config"
	/* 	"github.com/yvv4git/users-service/internal/servers"
	   	"google.golang.org/grpc" */)

const configPath = "config/main"

func main() {
	fmt.Println("Server")

	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal(err)
	}
	//config.setUpByViper()
	log.Println(cfg)

	// Create and register grpc server.
	/* 	grpcServer := grpc.NewServer()
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
	   	} */
}
