package main

import (
	"fmt"
)

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))                      // golang 1, ruby 2, js 4
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))                      // C 1, D 2, B 3, A 4
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))                                     // football 1, basketball 1, tenis 1
	fmt.Println(MostAppearItem([]string{"gajah", "burung", "ayam", "ayam", "gajah", "burung", "gajah", "ayam"})) // burung 2, ayam 3, gajah 3
}

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	mapping := make(map[string]int)
	var simpanNama []string
	var simpanJumlah []int

	for _, value := range items {
		var value2, sudahAda = mapping[value]

		if sudahAda {
			mapping[value] = value2 + 1
		} else {
			mapping[value] = value2 + 1
		}
	}

	for key, value := range mapping {
		simpanNama = append(simpanNama, key)
		simpanJumlah = append(simpanJumlah, value)
	}

	cek := false
	for !cek {
		cek = true
		for i := 0; i < len(simpanJumlah)-1; i++ {
			if simpanJumlah[i] > simpanJumlah[i+1] {
				temp := simpanJumlah[i]
				simpanJumlah[i] = simpanJumlah[i+1]
				simpanJumlah[i+1] = temp

				temp2 := simpanNama[i]
				simpanNama[i] = simpanNama[i+1]
				simpanNama[i+1] = temp2

				cek = false
			}
		}
	}

	pair1 := []pair{}
	for i := 0; i < len(simpanJumlah); i++ {
		pair2 := []pair{
			{simpanNama[i], simpanJumlah[i]},
		}
		pair1 = append(pair1, pair2...)
	}

	return pair1
}
