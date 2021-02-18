package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/yvv4git/users-service/internal/api"
	"github.com/yvv4git/users-service/internal/config"
	"github.com/yvv4git/users-service/internal/repository/sqlite"
	"github.com/yvv4git/users-service/internal/servers"
	"github.com/yvv4git/users-service/internal/services"
	"google.golang.org/grpc"
)

const configPath = "config/main"

func main() {
	log.Println("Server start")

	// Init config.
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal(err)
	}

	// Init db connection.
	db, err := sqlite.NewSqliteDB(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	// Init repository, service.
	usersRepo := sqlite.NewUsersRepository(db)
	usersService := services.NewUsersService(usersRepo)
	usersSrv := servers.NewUsersServer(usersService)

	// For graceful shutdown.
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Create and register grpc server.
	grpcServer := grpc.NewServer()
	api.RegisterUsersServer(grpcServer, usersSrv)

	// Bind port.
	address := fmt.Sprintf("%s:%s", "", strconv.Itoa(cfg.Server.Port))
	tcpListner, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	// Run grpc server.
	go func() {
		if err = grpcServer.Serve(tcpListner); err != nil {
			log.Fatal(err)
		}
	}()

	// Gracefull exit.
	<-exit
	log.Println("Stopping app...")
	grpcServer.Stop()

	if err := db.Close(); err != nil {
		log.Println("DB closed problems")
	}
}
