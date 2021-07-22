package main

import (
	"fmt"
	"go-test/api"
	"net/http"
)

func main() {
	fmt.Println("Start server")

	http.HandleFunc("/", api.Handler)

	http.ListenAndServe(":3000", nil)
}
