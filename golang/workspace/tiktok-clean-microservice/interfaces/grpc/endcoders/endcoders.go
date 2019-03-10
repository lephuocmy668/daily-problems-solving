package endcoders

import (
	"context"

	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/endpoints"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/interfaces/grpc/pb"
)

func EncodeUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoints.GetUserResponse)

	return &pb.UserResponse{
		Id:        res.User.ID.String(),
		FirstName: res.User.FirstName,
		LastName:  res.User.LastName,
	}, nil
}
