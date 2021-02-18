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
	//createUser(client)
	//readUser(client)
	//updateUser(client)
	delUser(client)
}

func createUser(client api.UsersClient) {
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

	log.Println(res.GetResult())
}

func readUser(client api.UsersClient) {
	res, err := client.Read(
		context.Background(),
		&api.ReadRequest{
			Id: 1,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetName(), res.GetEmail(), res.GetAge())
}

func updateUser(client api.UsersClient) {
	res, err := client.Update(
		context.Background(),
		&api.UpdateRequest{
			Id:    1,
			Name:  "Superman",
			Email: "super@gmail.ru",
			Age:   777,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetStatus())
}

func delUser(client api.UsersClient) {
	res, err := client.Del(
		context.Background(),
		&api.DelRequest{
			Id: 1,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetStatus())
}
