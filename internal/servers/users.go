package servers

import (
	"context"
	"log"

	"github.com/yvv4git/users-service/internal/api"
	"github.com/yvv4git/users-service/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Users ...
type Users struct {
	userService *services.UsersService
}

// Create ...
func (u *Users) Create(context.Context, *api.CreateRequest) (*api.CreateResponse, error) {
	//log.Println(u.userService)
	log.Println("Create")
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

// Read ...
func (u *Users) Read(context.Context, *api.ReadRequest) (*api.ReadResponse, error) {
	log.Println("Read")
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}

// Update ...
func (u *Users) Update(context.Context, *api.UpdateRequest) (*api.UpdateResponse, error) {
	log.Println("Update")
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

// Del ...
func (u *Users) Del(context.Context, *api.DelRequest) (*api.DelResponse, error) {
	log.Println("Del")
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}

// NewUsersServer is used as constructor for users server protocol.
func NewUsersServer(usersService *services.UsersService) *Users {
	return &Users{
		userService: usersService,
	}
}
