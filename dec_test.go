package stdlib

import (
	"math/big"
	"testing"
)

func Test_DEC(t *testing.T) {
	a := big.NewInt(37)
	t.Log(DECToBHex(a))

}
