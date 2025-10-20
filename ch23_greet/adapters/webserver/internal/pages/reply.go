package pages

import (
	"fmt"

	"github.com/go-rod/rod"
)

type Reply struct {
	Page *rod.Page
}

func (r *Reply) ReadReply() (string, error) {
	greeting, err := r.Page.Element("#reply")
	if err != nil {
		return "", fmt.Errorf("couldnt find #reply on Page")
	}
	return greeting.Text()
}
