package main

import (
	"context"
	"log"
	"net/http"

	"ch22/graceful/acceptancetests"
	"ch22/graceful/grace"
)

func main() {
	var (
		ctx        = context.Background()
		httpServer = &http.Server{Addr: ":8080", Handler: http.HandlerFunc(acceptancetests.SlowHandler)}
		server     = grace.NewServer(httpServer)
	)

	if err := server.ListenAndServe(ctx); err != nil {
		log.Fatalf("didnt shutdown gracefully, some responses may have been lost %v", err)
	}

	log.Println("shutdown gracefully! all responses were sent")
}
