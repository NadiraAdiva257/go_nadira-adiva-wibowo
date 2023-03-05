package main

import (
	"fmt"
)

func main() {
	fmt.Println(PlayingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}, {4, 3}}, []int{4, 3})) // {3, 4}
	fmt.Println(PlayingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}, {3, 6}}, []int{3, 6})) // {6, 5}
	fmt.Println(PlayingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1}))                 // gak ada
}

func PlayingDomino(cards [][]int, deck []int) interface{} {
	var hasil interface{}
	var kemungkinan [][]int
	ada := false

	for i := 0; i < len(cards); i++ {
		for j := 0; j < 2; j++ {
			if cards[i][j] == deck[0] || cards[i][j] == deck[1] {
				ada = true
				kemungkinan = append(kemungkinan, cards[i])
				break
			}
		}
	}

	var jumlah int
	var tampungJumlah []int

	if ada == true {
		fmt.Println("kemungkinan-kemungkinan:", kemungkinan)

		for i := 0; i < len(kemungkinan); i++ {
			for j := 0; j < 2; j++ {
				jumlah += kemungkinan[i][j]
			}

			tampungJumlah = append(tampungJumlah, jumlah)
			jumlah = 0
		}

		max := GetMax(tampungJumlah)
		for i := 0; i < len(kemungkinan); i++ {
			for j := 0; j < 2; j++ {
				jumlah += kemungkinan[i][j]
			}

			if max == jumlah {
				hasil = kemungkinan[i]
				break
			} else {
				jumlah = 0
			}
		}
	} else {
		hasil = "tutup kartu"
	}

	return hasil
}

func GetMax(jumlah []int) (max int) {
	max = jumlah[0]

	for _, value := range jumlah {
		if value > max {
			max = value
		}
	}

	return
}
