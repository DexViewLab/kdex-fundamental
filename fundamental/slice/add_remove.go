package slice

import "github.com/adam-hanna/arrayOperations"

// CalcAddAndRemove 算两个数组之间的diff
// 传入两个slice，old, new
// 传出两个slice，里面放着添加了哪些元素，删除了哪些元素
func CalcAddAndRemove(old, new interface{}) (add interface{}, remove interface{}) {
	// 先求交
	z, ok := arrayOperations.Intersect(old, new)
	if !ok {
		return
	}

	// add = diff(z, new)
	a, ok := arrayOperations.Difference(z.Interface(), new)
	if ok {
		add = a.Interface()
	}
	// remove = diff(z, old)
	r, ok := arrayOperations.Difference(z.Interface(), old)
	if ok {
		remove = r.Interface()
	}

	return
}
