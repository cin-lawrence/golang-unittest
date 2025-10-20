package main

import (
	"log"
	"net/http"

	"ch23/greet/adapters/httpserver"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", httpserver.NewHandler()))
}
