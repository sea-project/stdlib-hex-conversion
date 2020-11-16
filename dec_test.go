package hex

import "testing"

func Test_DEC(t *testing.T) {
	a := 37
	t.Log(DECToBIN(a))
	t.Log(DECToHEX(a))
	t.Log(DECToBHex(a))
}
