package main

import (
	"fmt"
	"net/http"

	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/common"
	"github.com/lephuocmy668/daily-problems-solving/golang/workspace/data-structs/linkedlist"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have visited %s in `%s`!", r.URL.Path, common.AppName)
}

func main() {
	lk := linkedlist.LinkedList{}
	lk.Append("11111")
	lk.Append("22222")
	// lk.Append("33333")
	// lk.Append("44444")
	// lk.DeleteByValue("11111")
	// lk.DeleteByValue("22222")
	// fmt.Println("=====1===", lk.Head)
	lk.DeleteTail()
	fmt.Println("=====1===", lk.Tail)
	common.StartServer("8080", handle)
}
