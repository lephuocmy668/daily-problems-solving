package main

import (
	"fmt"
	"net/http"

	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/core"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/data-structs/linkedlist"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have visited %s in `%s`!", r.URL.Path, core.AppName)
}

func main() {
	lk := linkedlist.LinkedList{}
	lk.Append(1)
	lk.Append(2)
	lk.Append(3)
	lk.Append(4)

	fmt.Println("=====1===", lk.Tail)
	fmt.Println("=====1===", lk.Head)

	core.StartServer("8080", handle)
}
