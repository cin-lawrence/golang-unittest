package main_test

import (
	"fmt"
	"testing"

	"ch23/greet/adapters"
	"ch23/greet/adapters/grpcserver"
	"ch23/greet/specs"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	var (
		port   = "50051"
		driver = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	t.Cleanup(driver.Close)
	adapters.StartDockerServer(t, port, "grpcserver")
	specs.AssertGreetSpecification(t, &driver)
	specs.AssertCurseSpecification(t, &driver)
}
