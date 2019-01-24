package decoders

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/endpoints"
)

// DecodeCreateUserRequest ...
func DecodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// DecodeUpdateUserRequest ...
func DecodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// DecodeGetUsersRequest ...
func DecodeGetUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoints.GetUsersRequest{}
	return req, nil
}

// DecodeGetUserRequest ...
func DecodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoints.GetUserRequest{}
	return req, nil
}
