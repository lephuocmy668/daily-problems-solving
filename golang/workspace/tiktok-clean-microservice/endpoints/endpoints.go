package endpoints

import (
	"github.com/go-kit/kit/endpoint"

	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/businesslogics/usecases"
)

// Endpoints ...
type Endpoints struct {
	CreateUser endpoint.Endpoint
	UpdateUser endpoint.Endpoint
	GetUsers   endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

// MakeEndpoints returns an Endpoints struct
func MakeEndpoints(usc usecases.UseCases) Endpoints {
	return Endpoints{
		CreateUser: MakeCreateUserEndpoint(usc),
		UpdateUser: MakeUpdateUserEndpoint(usc),
		GetUsers:   MakeGetUsersEndpoint(usc),
		GetUser:    MakeGetUserEndpoint(usc),
	}
}
