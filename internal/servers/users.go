package servers

import (
	"context"

	"github.com/yvv4git/users-service/internal/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Users ...
type Users struct{}

// Create ...
func (*Users) Create(context.Context, *api.CreateRequest) (*api.CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

// Read ...
func (*Users) Read(context.Context, *api.ReadRequest) (*api.ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}

// Update ...
func (*Users) Update(context.Context, *api.UpdateRequest) (*api.UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

// Del ...
func (*Users) Del(context.Context, *api.DelRequest) (*api.DelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}

// NewUsersServer is used as constructor for users server protocol.
func NewUsersServer() *Users {
	return &Users{}
}
