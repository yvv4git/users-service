package servers

import (
	"context"

	"github.com/yvv4git/users-service/internal/rpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Users ...
type Users struct{}

// Create ...
func (*Users) Create(context.Context, *rpc.CreateRequest) (*rpc.CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

// Read ...
func (*Users) Read(context.Context, *rpc.ReadRequest) (*rpc.ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}

// Update ...
func (*Users) Update(context.Context, *rpc.UpdateRequest) (*rpc.UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

// Del ...
func (*Users) Del(context.Context, *rpc.DelRequest) (*rpc.DelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
