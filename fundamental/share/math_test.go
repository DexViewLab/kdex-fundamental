package share

import (
	"strconv"
	"testing"
	"time"
)

func Test_RandFloat(t *testing.T) {
	for i := 0; i < 10; i ++ {
		flo := RandFloat(-1, 1)
		f := strconv.FormatFloat(flo, 'f', 8, 64)
		println(f)
		time.Sleep(1 * time.Second)
	}
}