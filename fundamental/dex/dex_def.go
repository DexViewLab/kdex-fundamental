package dex

import "strings"

type Dex int32

const (
	Dex_uniswap Dex = 0
	Dex_sushi   Dex = 1
	Dex_pancake Dex = 2
	Dex_UNKNOWN Dex = 999
)

var Dex_name = map[int32]string{
	0:   "uniswap",
	1:   "sushi",
	2:   "pancake",
	999: "UNKNOWN",
}
var Dex_value = map[string]int32{
	"uniswap": 0,
	"sushi":   1,
	"pancake": 2,
	"UNKNOWN": 999,
}

func StringToDex(str string) Dex {
	d, ok := Dex_value[str]
	if ok {
		return Dex(d)
	}
	d, ok = Dex_value[strings.ToUpper(str)]
	if ok {
		return Dex(d)
	}
	d, ok = Dex_value[strings.ToLower(str)]
	if ok {
		return Dex(d)
	}
	return Dex_UNKNOWN
}
