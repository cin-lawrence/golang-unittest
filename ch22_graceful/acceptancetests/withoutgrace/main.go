package main

import (
	"log"
	"net/http"

	"ch22/graceful/acceptancetests"
)

func main() {
	server := &http.Server{Addr: ":8081", Handler: http.HandlerFunc(acceptancetests.SlowHandler)}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
