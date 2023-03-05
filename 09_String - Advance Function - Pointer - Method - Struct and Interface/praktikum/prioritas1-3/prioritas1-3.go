package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
	fmt.Println(Compare("ILALANG", "ICA"))    // gak ada
	fmt.Println(Compare("ND", "PANDA"))       // ND
}

func Compare(a, b string) string {
	var hasil string

	if strings.Contains(a, b) {
		hasil = b
	} else if strings.Contains(b, a) {
		hasil = a
	} else {
		hasil = "tidak ada kata yang menjadi bagian dari kata lain"
	}

	return hasil
}
