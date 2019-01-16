package common

import (
	"fmt"
	"log"
	"net/http"
)

const AppName = "app"
const AnotherAppName = "another-app"

func StartServer(port string, handlerFunc http.HandlerFunc) {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting web server on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
