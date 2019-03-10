package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"

	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/businesslogics/domains"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/businesslogics/repositories"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/businesslogics/usecases"
)

// CreateUserRequest is struct which describe create user request
type CreateUserRequest struct {
	domains.User
}

// CreateUserResponse is struct which describe create user response
type CreateUserResponse struct {
	domains.User
}

// MakeCreateUserEndpoint return execute create user endpoint
func MakeCreateUserEndpoint(usc usecases.UseCases) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)

		user, err := usc.User.Create(&req.User)
		if err != nil {
			return nil, err
		}

		return CreateUserResponse{*user}, err
	}
}

// UpdateUserRequest is struct which describe update user request
type UpdateUserRequest struct {
	domains.User
}

// UpdateUserResponse is struct which describe update user response
type UpdateUserResponse struct {
	domains.User
}

// MakeUpdateUserEndpoint return execute update user endpoint
func MakeUpdateUserEndpoint(usc usecases.UseCases) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		user, err := usc.User.Update(&req.User)
		if err != nil {
			return nil, err
		}
		return UpdateUserResponse{*user}, err
	}
}

// GetUsersRequest is struct which describe get users request
type GetUsersRequest struct {
	repositories.UserRequest
}

// GetUsersResponse is struct which describe get users response
type GetUsersResponse struct {
	users []domains.User
}

// MakeGetUsersEndpoint return execute get users endpoint
func MakeGetUsersEndpoint(usc usecases.UseCases) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUsersRequest)
		users, err := usc.User.Find(req.UserRequest)
		if err != nil {
			return nil, err
		}
		return GetUsersResponse{users}, err
	}
}

// GetUserRequest is struct which describe get user request
type GetUserRequest struct {
	repositories.UserRequest
}

// GetUserResponse is struct which describe get user response
type GetUserResponse struct {
	User domains.User
}

// MakeGetUserEndpoint return execute get user endpoint
func MakeGetUserEndpoint(usc usecases.UseCases) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		user, err := usc.User.FindOne(req.UserRequest)
		if err != nil {
			return nil, err
		}
		if user == nil {
			return nil, errors.New("User not found")
		}
		return GetUserResponse{
			User: *user,
		}, err
	}
}
