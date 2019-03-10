package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"

	tiktokConfig "github.com/lephuocmy668/daily-problems-solving/golang/workspace/core/config"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/businesslogics/usecases"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/endpoints"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/interfaces/grpc/pb"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/interfaces/grpc/server"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/interfaces/searchengine"
)

func StartGRPCServer() {
	// get config
	var config tiktokConfig.Config
	err := envconfig.Process("tiktok", &config)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println(config)

	// infrastructure
	userIndex := searchengine.UserIndex{}
	useCases := usecases.MakeUseCases(&userIndex)

	// start
	ctx := context.Background()
	errors := make(chan error)

	go func() {
		listener, err := net.Listen("tcp", ":9090")
		if err != nil {
			errors <- err
			return
		}

		gRPCServer := grpc.NewServer()
		pb.RegisterUserServer(gRPCServer, server.NewGRPCServer(ctx, endpoints.MakeEndpoints(*useCases)))

		fmt.Println("gRPC listen on 9090")
		errors <- gRPCServer.Serve(listener)
	}()

	fmt.Println(<-errors)
}
