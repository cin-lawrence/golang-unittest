package iteration

import "strings"

const repeatCount = 5

func Repeat(char string, count int) string {
	if count < 0 {
		count = repeatCount
	}
	var repeated strings.Builder
	for i := 0; i < count; i++ {
		repeated.WriteString(char)
	}
	return repeated.String()
}
