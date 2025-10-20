package main

import (
	"testing"
	"time"

	"ch22/graceful/acceptancetests"
	"ch22/graceful/assert"
)

const (
	port = "8080"
	url  = "http://localhost:" + port
)

func TestGracefulShutdown(t *testing.T) {
	cleanup, sendInterrupt, err := acceptancetests.LaunchTestProgram(port)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(cleanup)
	assert.CanGet(t, url)

	time.AfterFunc(50*time.Millisecond, func() {
		assert.NoError(t, sendInterrupt())
	})
	assert.CanGet(t, url)

	assert.CantGet(t, url)
}
