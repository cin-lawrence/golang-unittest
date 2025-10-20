package main

import (
	"testing"
	"time"

	"ch22/graceful/acceptancetests"
	"ch22/graceful/assert"
)

const (
	port = "8081"
	url  = "http://localhost:" + port
)

func TestNonGracefulShutdown(t *testing.T) {
	cleanup, sendInterrupt, err := acceptancetests.LaunchTestProgram(port)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(cleanup)
	assert.CanGet(t, url)

	time.AfterFunc(50*time.Millisecond, func() {
		assert.NoError(t, sendInterrupt())
	})
	assert.CantGet(t, url)

	assert.CantGet(t, url)
}
