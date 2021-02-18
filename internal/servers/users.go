package servers

import (
	"context"

	"github.com/yvv4git/users-service/internal/api"
	"github.com/yvv4git/users-service/internal/domain"
	"github.com/yvv4git/users-service/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Users ...
type Users struct {
	userService *services.UsersService
}

// Create users.
func (u *Users) Create(ctx context.Context, request *api.CreateRequest) (*api.CreateResponse, error) {
	var user domain.Users
	user.Name = request.Name
	user.Email = request.Email
	user.Age = request.Age

	_, err := u.userService.Create(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &api.CreateResponse{Result: true}, nil
}

// Read user.
func (u *Users) Read(ctx context.Context, request *api.ReadRequest) (*api.ReadResponse, error) {
	user, err := u.userService.Read(request.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &api.ReadResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
	}, nil
}

// Update user.
func (u *Users) Update(ctx context.Context, request *api.UpdateRequest) (*api.UpdateResponse, error) {
	var user domain.Users
	user.ID = request.Id
	user.Name = request.Name
	user.Email = request.Email
	user.Age = request.Age

	err := u.userService.Update(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &api.UpdateResponse{
		Status: true,
	}, nil
}

// Del user.
func (u *Users) Del(ctx context.Context, request *api.DelRequest) (*api.DelResponse, error) {
	err := u.userService.Del(request.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &api.DelResponse{
		Status: true,
	}, nil
}

// NewUsersServer is used as constructor for users server protocol.
func NewUsersServer(usersService *services.UsersService) *Users {
	return &Users{
		userService: usersService,
	}
}
