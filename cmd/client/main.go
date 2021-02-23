package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/yvv4git/users-service/internal/api"
	"google.golang.org/grpc"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app  = kingpin.New("usersClient", "A users client on golang.")
	host = app.Flag("host", "IP address of server.").Required().IP()
	port = app.Flag("port", "Port of server.").Required().String()

	create      = app.Command("create", "Add user to system.")
	createName  = create.Arg("name", "User name.").Required().String()
	createEmail = create.Arg("email", "User email.").Required().String()
	createAge   = create.Arg("age", "User age").Required().Int32()

	read      = app.Command("read", "Find user in system.")
	readID    = read.Arg("id", "User id.").Required().Int64()
	readName  = read.Arg("name", "User name.").String()
	readEmail = read.Arg("email", "User email.").String()
	readAge   = read.Arg("age", "User age").String()

	update      = app.Command("update", "Update user params.")
	updateID    = update.Arg("user_id", "User id.").Required().Int64()
	updateName  = update.Arg("name", "User name.").String()
	updateEmail = update.Arg("email", "User email.").String()
	updateAge   = update.Arg("age", "User age").Int32()

	delete   = app.Command("delete", "Delete user from system.")
	deleteID = delete.Arg("id", "User id.").Required().Int64()
)

func main() {
	var funcCallBack func(client api.UsersClient)

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case create.FullCommand():
		funcCallBack = createUser
	case read.FullCommand():
		funcCallBack = readUser
	case update.FullCommand():
		funcCallBack = updateUser
	case delete.FullCommand():
		funcCallBack = delUser
	}

	serverAddress := fmt.Sprintf("%s:%s", *host, *port)
	log.Println("Server address: ", serverAddress)

	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := api.NewUsersClient(conn)
	funcCallBack(client)
}

func createUser(client api.UsersClient) {
	log.Println("Add command.")

	res, err := client.Create(
		context.Background(),
		&api.CreateRequest{
			Name:  *createName,
			Email: *createEmail,
			Age:   *createAge,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetResult())
}

func readUser(client api.UsersClient) {
	log.Println("Read command.")

	res, err := client.Read(
		context.Background(),
		&api.ReadRequest{
			Id: *readID,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetName(), res.GetEmail(), res.GetAge())
}

func updateUser(client api.UsersClient) {
	log.Println("Update command.")

	res, err := client.Update(
		context.Background(),
		&api.UpdateRequest{
			Id:    *updateID,
			Name:  *updateName,
			Email: *updateEmail,
			Age:   *updateAge,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetStatus())
}

func delUser(client api.UsersClient) {
	log.Println("Delete command.")

	res, err := client.Del(
		context.Background(),
		&api.DelRequest{
			Id: *deleteID,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetStatus())
}
