package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt64SliceToInterfaceSlice(t *testing.T) {
	a := []int64{1, 2, 3}
	b := Int64SliceToInterfaceSlice(a)

	assert.Equal(t, len(b), 3)
	for i, v := range b {
		assert.Equal(t, a[i], v.(int64))
	}
}

func TestStringSliceToInterfaceSlice(t *testing.T) {
	a := []string{"11", "22", "33"}
	b := StringSliceToInterfaceSlice(a)

	assert.Equal(t, len(b), 3)
	for i, v := range b {
		assert.Equal(t, a[i], v.(string))
	}
}
