package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayToString(t *testing.T) {
	assert.Equal(t, ArrayToString([]int{1, 2, 3, 4}), "1,2,3,4")
	assert.Equal(t, ArrayToString([]int{1}), "1")
	assert.Equal(t, ArrayToString([]int{}), "")

	s, _ := StringToIntSlice("1,2,3,4")
	assert.Equal(t, s, []int{1, 2, 3, 4})
}

func TestStringArrayToString(t *testing.T) {
	// []string != ""
	assert.Equal(t, "a,b,c", StringArrayToString([]string{"a", "b", "c"}))
	// []string == ""
	assert.Equal(t, "", StringArrayToString([]string{}))
	// []string == 1
	assert.Equal(t, "a", StringArrayToString([]string{"a"}))
}
