package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/endpoints"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/interfaces/httpserver/decoders"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/interfaces/httpserver/presenters"
)

// NewHTTPHandler ...
func NewHTTPHandler(endpoints endpoints.Endpoints, logger log.Logger, useCORS bool) http.Handler {
	r := chi.NewRouter()

	if useCORS {
		cors := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowCredentials: true,
		})
		r.Use(cors.Handler)
	}

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(presenters.EncodeError),
	}

	r.Route("/users", func(r chi.Router) {
		r.Post("/", httptransport.NewServer(
			endpoints.CreateUser,
			decoders.DecodeCreateUserRequest,
			presenters.EncodeResponse,
			options...,
		).ServeHTTP)

		r.Get("/", httptransport.NewServer(
			endpoints.GetUsers,
			decoders.DecodeGetUsersRequest,
			presenters.EncodeResponse,
			options...,
		).ServeHTTP)

		r.Get("/{id}", httptransport.NewServer(
			endpoints.GetUser,
			decoders.DecodeGetUserRequest,
			presenters.EncodeResponse,
			options...,
		).ServeHTTP)

		r.Put("/{id}", httptransport.NewServer(
			endpoints.UpdateUser,
			decoders.DecodeUpdateUserRequest,
			presenters.EncodeResponse,
			options...,
		).ServeHTTP)
	})

	return r
}
