package test8

func HitungPeluangKemunculanDadu() int {
	sisiDadu := 6
	kemunculan := 0
	for dadu1 := 1; dadu1 <= sisiDadu; dadu1++ {
		for dadu2 := 1; dadu2 <= sisiDadu; dadu2++ {
			for dadu3 := 1; dadu3 <= sisiDadu; dadu3++ {
				jumlahSisiDadu := dadu1 + dadu2 + dadu3
				if jumlahSisiDadu > 2 && jumlahSisiDadu < 19 {
					kemunculan++
				}
			}
		}
	}
	return kemunculan
}
