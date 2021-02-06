package main

import (
	"context"
	"fmt"
	"log"

	"github.com/yvv4git/users-service/internal/api"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client")

	// Connect to grpc server.
	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	// Create client with connection.
	client := api.NewUsersClient(conn)

	// Use client method Create for example.
	res, err := client.Create(
		context.Background(),
		&api.CreateRequest{
			Name:  "Vladimir",
			Email: "yvv4recon@gmail.com",
			Age:   32,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.GetResult())
}
