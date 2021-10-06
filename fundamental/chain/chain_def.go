package chain

import "strings"

type Chain int32

const (
	Chain_eth Chain = 0
	Chain_bsc   Chain = 1
	Chain_UNKNOWN Chain = 999
)

var Chain_name = map[int32]string{
	0:   "eth",
	1:   "bsc",
	999: "UNKNOWN",
}
var Chain_value = map[string]int32{
	"eth": 0,
	"bsc":   1,
	"UNKNOWN": 999,
}

func StringToChain(str string) Chain {
	c, ok := Chain_value[str]
	if ok {
		return Chain(c)
	}
	c, ok = Chain_value[strings.ToUpper(str)]
	if ok {
		return Chain(c)
	}
	c, ok = Chain_value[strings.ToLower(str)]
	if ok {
		return Chain(c)
	}
	return Chain_UNKNOWN
}
