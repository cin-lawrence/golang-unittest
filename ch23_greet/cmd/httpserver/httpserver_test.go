package main_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"ch23/greet/adapters"
	"ch23/greet/adapters/httpserver"
	"ch23/greet/specs"
)

func TestHttpServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	var (
		port   = "8080"
		driver = httpserver.Driver{
			BaseURL: fmt.Sprintf("http://localhost:%s", port),
			Client: &http.Client{
				Timeout: 1 * time.Second,
			},
		}
	)

	adapters.StartDockerServer(t, port, "httpserver")
	specs.AssertGreetSpecification(t, &driver)
	specs.AssertCurseSpecification(t, &driver)
}
