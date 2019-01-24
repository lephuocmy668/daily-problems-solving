package core

import (
	"fmt"
	"log"
	"net/http"
)

const AppName = "app"
const AnotherAppName = "tiktok-clean-microservice"

func StartServer(port string, handlerFunc http.HandlerFunc) {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting web server on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
