package main

import (
	"example/main/api"
	"net/http"
)

func main() {
	// Code
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)

}
