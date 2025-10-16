package foo

import "lawrence/sample/bar"

func Something(a, b int) int {
	return bar.Something(a, b)
}
