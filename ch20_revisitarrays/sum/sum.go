package sum

import (
	"golang-unittest/ch20/util"
)

func Sum(numbers []int) int {
	return util.Reduce(numbers, func(acc, x int) int {
		return acc + x
	}, 0)
}

func SumAllTails(numbers ...[]int) []int {
	return util.Reduce(numbers, func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		}

		tail := x[1:]
		return append(acc, Sum(tail))
	}, []int{})
}
