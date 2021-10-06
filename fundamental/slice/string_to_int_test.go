package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSliceToIntSlice(t *testing.T) {
	a := []string{"1", "2", "3", "4"}
	b, err := StringSliceToIntSlice(a)
	assert.Nil(t, err)
	assert.Equal(t, 1, b[0])
	assert.Equal(t, 2, b[1])
	assert.Equal(t, 3, b[2])
	assert.Equal(t, 4, b[3])

}

func TestStringSliceToInt64Slice(t *testing.T) {
	a := []string{"1", "2", "3", "4"}
	b, err := StringSliceToInt64Slice(a)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), b[0])
	assert.Equal(t, int64(2), b[1])
	assert.Equal(t, int64(3), b[2])
	assert.Equal(t, int64(4), b[3])

}
