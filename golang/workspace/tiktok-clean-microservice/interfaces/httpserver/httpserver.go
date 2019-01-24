package httpserver

import (
	"log"
	"net/http"
	"os"

	gokitlog "github.com/go-kit/kit/log"
	"github.com/kelseyhightower/envconfig"

	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/core/httpserver"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/businesslogics/usecases"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/endpoints"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/interfaces/httpserver/controllers"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/tiktok-clean-microservice/interfaces/searchengine"
	tiktokConfig "github.com/lephuocmy668/tiktok/tiktok/golang/product/config"
)

const useCORS = false

func RunHttpServer() {
	var logger gokitlog.Logger
	{
		logger = gokitlog.NewLogfmtLogger(gokitlog.NewSyncWriter(os.Stdout))
		logger = gokitlog.With(logger, "ts", gokitlog.DefaultTimestampUTC)
		logger = gokitlog.With(logger, "caller", gokitlog.DefaultCaller)
	}

	var config tiktokConfig.Config
	err := envconfig.Process("tiktok", &config)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println(config)

	userIndex := searchengine.UserIndex{}

	useCases := usecases.MakeUseCases(&userIndex)

	var h http.Handler
	{
		h = controllers.NewHTTPHandler(
			endpoints.MakeEndpoints(*useCases),
			logger,
			useCORS,
		)
	}
	err = httpserver.InitHTTPServer(h, config.HTTPServiceConfigBindIP+":"+config.HTTPServiceConfigPort)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return
}
