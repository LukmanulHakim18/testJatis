package test4

import (
	"fmt"
	"testing"
)

func TestHitungKarakter(t *testing.T) {
	// t.Skip("tidak urut")
	kata := "dani Maulana"
	hasil := "d4a2nimul"

	res := HitungKarakter(kata)
	fmt.Println(hasil)
	if res != hasil {
		// t.Error("not same", res, hasil)
	}
}
