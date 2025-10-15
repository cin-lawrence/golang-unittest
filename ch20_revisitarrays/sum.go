package main

func Sum(numbers []int) int {
	return Reduce(numbers, func(acc, x int) int {
		return acc + x
	}, 0)
}

func SumAllTails(numbers ...[]int) []int {
	return Reduce(numbers, func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		}

		tail := x[1:]
		return append(acc, Sum(tail))
	}, []int{})
}
