package slice

import (
	"strconv"
	"strings"
)

// StringToIntSlice "1,2,3" -> {1, 2, 3}
func StringToIntSlice(str string) ([]int, error) {
	return StringSliceToIntSlice(strings.Split(str, ","))
}

// StringToInt64Slice "1,2,3" -> {1, 2, 3}
func StringToInt64Slice(str string) ([]int64, error) {
	return StringSliceToInt64Slice(strings.Split(str, ","))
}

// StringSliceToIntSlice {"1", "2", "3"} -> {1, 2, 3}
func StringSliceToIntSlice(str []string) ([]int, error) {
	v := []int{}
	for _, s := range str {
		i, err := strconv.Atoi(s)
		if err == nil {
			v = append(v, i)
		} else {
			return nil, err
		}
	}
	return v, nil
}

// StringSliceToInt64Slice {"1", "2", "3"} -> {1, 2, 3}
func StringSliceToInt64Slice(str []string) ([]int64, error) {
	v := []int64{}
	for _, s := range str {
		i, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			v = append(v, i)
		} else {
			return nil, err
		}
	}
	return v, nil
}
