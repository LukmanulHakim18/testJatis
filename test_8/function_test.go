package test8

import (
	"fmt"
	"testing"
)

func TestHitungKarakter(t *testing.T) {
	// t.Skip("tidak urut")
	/*sebenarnya total kemunculan dadu adaldah seluruh totl peluang
	karna jumlah dari 3 sisi dadu terkecilnya adalh 3 dan terbesarnya 18
	ini menjadikan rumus 6pangkat3 =216*/
	ekspektasi := 216
	res := HitungPeluangKemunculanDadu()

	fmt.Println("jumlah kemunculan :", res)
	if res != ekspektasi {
		t.Error("hasil salah")
	}

}
