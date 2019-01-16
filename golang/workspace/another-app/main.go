package main

import (
	"fmt"
	"net/http"

	"github.com/lephuocmy668/daily-problem-solving/golang/workspace/common"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have visited %s in `%s`!", r.URL.Path, common.AnotherAppName)
}

func main() {
	common.StartServer("8181", handle)
}
