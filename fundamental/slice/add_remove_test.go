package slice

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddRemove(t *testing.T) {
	tester := func(old, new, add, remove interface{}) {
		a, r := CalcAddAndRemove(old, new)
		sortedAdd := add.([]int)
		sort.Ints(sortedAdd)
		sortedA := a.([]int)
		sort.Ints(sortedA)
		assert.Equal(t, sortedAdd, sortedA)

		sortedRemove := remove.([]int)
		sort.Ints(sortedRemove)
		sortedR := r.([]int)
		sort.Ints(sortedR)
		assert.Equal(t, sortedRemove, sortedR)
	}

	empty := []int{}
	one := []int{1}

	// add 1
	tester([]int{}, one, one, empty)

	// add ...
	tester([]int{}, []int{1, 2, 3, 4}, []int{1, 2, 3, 4}, empty)
	tester([]int{2, 3}, []int{1, 2, 3, 4}, []int{1, 4}, empty)

	// remove 1
	tester(one, empty, empty, one)
	tester([]int{1, 2, 3, 4}, one, empty, []int{2, 3, 4})

	// both
	tester([]int{1, 3}, []int{2, 3}, []int{2}, []int{1})

	// emtpy
	tester(empty, empty, empty, empty)

}
