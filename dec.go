package stdlib

import (
	"math/big"
)

var (
	Big0        = big.NewInt(0)
	Big36       = big.NewInt(36)
	//Base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func DECToBHex(i *big.Int) string {
	var chars []rune
	x := new(big.Int)
	for {
		x.Mod(i, Big36)
		chars = append(chars, rune(Base62Chars[x.Uint64()]))
		i.Div(i, Big36)
		if i.Cmp(Big0) == 0 {
			break
		}
	}
	// reverse slice
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
