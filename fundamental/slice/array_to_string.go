package slice

import (
	"fmt"
	"strings"
)

// ArrayToString []int{1,2,3,4} => "1,2,3,4"
func ArrayToString(a []int) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), ","), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}

// Array64ToString []int{1,2,3,4} => "1,2,3,4"
func Array64ToString(a []int64) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), ","), "[]")
}

// StringArrayToString []int{"a", "b", "c"} => "a,b,c"
func StringArrayToString(a []string) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), ","), "[]")
}
