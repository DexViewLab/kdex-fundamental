package slice

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/golang/glog"
)

type filtertype func(interface{}) bool

// InSlice checks given string in string slice or not.
func InSlice(v string, sl []string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// InSliceIface checks given interface in interface slice.
func InSliceIface(v interface{}, sl []interface{}) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// SliceRandList generate an int slice from min to max.
func SliceRandList(min, max int) []int {
	if max < min {
		min, max = max, min
	}
	length := max - min + 1
	t0 := time.Now()
	rand.Seed(int64(t0.Nanosecond()))
	list := rand.Perm(length)
	for index := range list {
		list[index] += min
	}
	return list
}

// SliceFilter generates a new slice after filter function.
func SliceFilter(slice []interface{}, a filtertype) (ftslice []interface{}) {
	for _, v := range slice {
		if a(v) {
			ftslice = append(ftslice, v)
		}
	}
	return
}

func checkSliceType(i interface{}) reflect.Value {
	p := reflect.ValueOf(i)
	for {
		switch p.Kind() {
		case reflect.Slice:
			return p
		case reflect.Ptr:
			p = p.Elem()
		default:
			glog.FatalDepth(0, "unknown type", p)
		}
	}
}

// SliceUnique 用于slice去重，填入slice地址
func SliceUnique(t interface{}) {
	newT := checkSliceType(t)
	length := newT.Len()

	sliceM := make(map[interface{}]struct{})
	for i := 0; i < length; i++ {
		sliceM[newT.Index(i).Interface()] = struct{}{}
	}

	i := 0
	for k := range sliceM {
		newT.Index(i).Set(reflect.ValueOf(k))
		i++
	}
	newT.SetLen(len(sliceM))
}
