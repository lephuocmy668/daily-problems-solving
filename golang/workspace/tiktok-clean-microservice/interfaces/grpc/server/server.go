package server

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/endpoints"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/interfaces/grpc/decoders"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/interfaces/grpc/endcoders"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/interfaces/grpc/pb"
)

type GRPCServer struct {
	getUser    grpctransport.Handler
	createUser grpctransport.Handler
}

func (s *GRPCServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	_, resp, err := s.getUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.UserResponse), nil
}

func (s *GRPCServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	_, resp, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.UserResponse), nil
}

func NewGRPCServer(_ context.Context, endpoint endpoints.Endpoints) pb.UserServer {
	return &GRPCServer{
		getUser: grpctransport.NewServer(
			endpoint.GetUser,
			decoders.DecodeGetUserRequest,
			endcoders.EncodeUserResponse,
		),
		createUser: grpctransport.NewServer(
			endpoint.CreateUser,
			decoders.DecodeCreateUserRequest,
			endcoders.EncodeUserResponse,
		),
	}
}
