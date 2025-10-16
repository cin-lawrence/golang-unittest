package sum_test

import (
	"reflect"
	"strings"
	"testing"

	"golang-unittest/ch20/assert"
	"golang-unittest/ch20/sum"
	"golang-unittest/ch20/util"
)

func TestSum(t *testing.T) {
	t.Run("collections of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := sum.Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := sum.SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := sum.SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		assert.AssertEqual(t, util.Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		assert.AssertEqual(t, util.Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := util.Find(numbers, func(x int) bool {
			return x%2 == 0
		})

		assert.AssertTrue(t, found)
		assert.AssertEqual(t, firstEvenNumber, 2)
	})

	type Person struct {
		Name string
	}

	t.Run("find the best programmer", func(t *testing.T) {
		people := []Person{
			Person{Name: "Kent Beck"},
			Person{Name: "Martin Fowler"},
			Person{Name: "Chris James"},
		}

		king, found := util.Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Chris")
		})

		assert.AssertTrue(t, found)
		assert.AssertEqual(t, king, Person{Name: "Chris James"})
	})
}
