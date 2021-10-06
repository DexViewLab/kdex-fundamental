package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInSlice(t *testing.T) {
	sl := []string{"A", "b"}
	if !InSlice("A", sl) {
		t.Error("should be true")
	}
	if InSlice("B", sl) {
		t.Error("should be false")
	}
}

func TestInSliceIface(t *testing.T) {
	sl := []string{"A", "b"}

	var interfaceSlice = make([]interface{}, len(sl))
	for i, d := range sl {
		interfaceSlice[i] = d
	}

	if !InSliceIface("A", interfaceSlice) {
		t.Error("should be true")
	}
	if InSliceIface("B", interfaceSlice) {
		t.Error("should be false")
	}
}

func TestSliceRandList(t *testing.T) {
	result := SliceRandList(1, 10)
	assert.Equal(t, 10, len(result))
	for _, value := range result {
		if value < 1 || value > 10 {
			assert.New(t)
		}
	}
}

func TestSliceFilter(t *testing.T) {
	sl := SliceRandList(1, 10)
	assert.Equal(t, 10, len(sl))
	num := 0
	for _, value := range sl {
		if value > 5 {
			num++
		}
	}
	fmt.Println(sl)

	var interfaceSlice = make([]interface{}, len(sl))
	for i, d := range sl {
		interfaceSlice[i] = d
	}

	newSlice := SliceFilter(interfaceSlice, func(value interface{}) bool {
		return value.(int) > 5
	})
	fmt.Println(newSlice)
	assert.Equal(t, num, len(newSlice))
}

func TestSliceUnique(t *testing.T) {
	var b = &[]string{"bbb", "aaa", "ccc", "aaa", "ccc"}
	SliceUnique(b)
	fmt.Println(b)
	assert.Equal(t, 3, len(*b))

	var a = &[]int{1, 2, 1, 1, 1, 1, 1, 2}
	SliceUnique(a)
	fmt.Println(a)
	assert.Equal(t, 2, len(*a))

}
