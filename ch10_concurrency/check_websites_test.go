package concurrency

import (
	"reflect"
	"testing"
)

const wrongWebsite = "waat://furhurterwe.geds"

func mockWebsiteChecker(url string) bool {
	return url != wrongWebsite
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		wrongWebsite,
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		wrongWebsite:                 false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
