package test4

import (
	"fmt"
	"sort"
	"strings"
)

func HitungKarakter(kata string) string {
	groupCharacter := make(map[string]int)
	newKata := strings.ReplaceAll(kata, " ", "")
	newKata = strings.ToLower(newKata)
	//fungsi untuk menghitung karakter
	for _, v := range newKata {
		groupCharacter[string(v)] += 1
	}
	keys := make([]string, 0, len(groupCharacter))
	for k := range groupCharacter {
		keys = append(keys, k)
	}
	sort.Strings(keys) //mengurutkan keyss

	var otuput string 
	for _, v := range keys {
		otuput = fmt.Sprintf("%s, %s:%v", otuput, v, groupCharacter[v])

	}
	fmt.Println("input :", kata)
	fmt.Println("output :", otuput)
	return otuput
}
