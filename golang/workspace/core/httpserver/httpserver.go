package httpserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// InitHTTPServer start a http server
func InitHTTPServer(httpHandler http.Handler, address string) error {
	stopServerListenerChannel := make(chan os.Signal)
	signal.Notify(stopServerListenerChannel, os.Interrupt)

	httpServer := &http.Server{
		Addr:    address,
		Handler: httpHandler,
	}

	go startServer(httpServer)
	<-stopServerListenerChannel
	log.Println("HTTP server is shutting down....")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	return httpServer.Shutdown(ctx)
}

func startServer(s *http.Server) {
	err := s.ListenAndServe()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
