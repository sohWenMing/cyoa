package main

import (
	"fmt"
	"net/http"

	"github.com/sohWenMing/cyoa/server"
)

func main() {
	mux := server.InitServer()
	fmt.Println("listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
