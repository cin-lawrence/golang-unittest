package main

func Find[A any](items []A, predicate func(A) bool) (value A, found bool) {
	for _, item := range items {
		if predicate(item) {
			return item, true
		}
	}
	return
}

func Reduce[A, B any](items []A, f func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, item := range items {
		result = f(result, item)
	}
	return result
}
