package decoders

import (
	"context"

	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/endpoints"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/interfaces/grpc/pb"
)

// DecodeCreateUserRequest ...
func DecodeCreateUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	gRPCRequest := r.(*pb.CreateUserRequest)

	req := endpoints.CreateUserRequest{}
	req.LastName = gRPCRequest.LastName
	req.FirstName = gRPCRequest.FirstName

	return req, nil
}

// DecodeGetUserRequest ...
func DecodeGetUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	gRPCRequest := r.(*pb.GetUserRequest)

	req := endpoints.GetUserRequest{}
	req.ID = gRPCRequest.Id

	return req, nil
}
