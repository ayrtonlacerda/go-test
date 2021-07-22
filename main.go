package main

import (
	"fmt"
	"go-test/api"
	"net/http"
)

func main() {
	fmt.Println("Start server")

	http.HandleFunc("/test", api.Handler)
	http.HandleFunc("/test2", api.Handler2)

	http.ListenAndServe(":3000", nil)
}
