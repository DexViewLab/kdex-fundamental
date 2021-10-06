package coin

import (
	"fmt"
	"strings"
)

// SymbolPairStringer ..
type SymbolPairStringer interface {
	SymbolPairString(base, baseAddress, quote string) string
}

// SymbolPair 一个币对，交易对
type SymbolPair struct {
	Base        string // 交易币
	BaseAddress string // 交易币地址
	Quote       string // 计价币
}

func (s *SymbolPair) String(stringer SymbolPairStringer) string {
	return stringer.SymbolPairString(s.Base, s.BaseAddress, s.Quote)
}

// StandardString .
func (s *SymbolPair) StandardString() string {
	return fmt.Sprintf("%s_%s_%s", s.Base, s.BaseAddress, s.Quote)
}

// CreateSymbolPair ..
func CreateSymbolPair(base, baseAddress, quote string) *SymbolPair {
	return &SymbolPair{
		Base:  strings.ToUpper(base),
		Quote: strings.ToUpper(quote),
		BaseAddress: baseAddress,
	}
}

// SymbolPairQuoteUSDT ..
func SymbolPairQuoteUSDT(base, baseAddress string) *SymbolPair {
	return CreateSymbolPair(base, baseAddress, "USDT")
}
