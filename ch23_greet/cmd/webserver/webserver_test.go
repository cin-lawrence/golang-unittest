package main_test

import (
	"fmt"
	"testing"

	"ch23/greet/adapters"
	"ch23/greet/adapters/webserver"
	"ch23/greet/specs"
	"github.com/alecthomas/assert/v2"
)

func TestGreeterWeb(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	var (
		port            = "8081"
		driver, cleanup = webserver.NewDriver(fmt.Sprintf("http://localhost:%s", port))
	)

	t.Cleanup(func() {
		assert.NoError(t, cleanup())
	})

	adapters.StartDockerServer(t, port, "webserver")
	specs.AssertGreetSpecification(t, driver)
	specs.AssertCurseSpecification(t, driver)
}
