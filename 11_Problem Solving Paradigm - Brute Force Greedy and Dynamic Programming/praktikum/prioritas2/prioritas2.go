package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))             // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50, 20})) // 40
}

func Frog(jumps []int) int {
	return Rekursif(len(jumps)-1, jumps)
}

func Rekursif(index int, jumps []int) int {
	var lompat1, lompat2 int

	if index == 0 {
		return 0
	} else {
		lompat1 = Rekursif(index-1, jumps) + int(math.Abs(float64(jumps[index-1])-float64(jumps[index])))
		if index > 1 {
			lompat2 = Rekursif(index-2, jumps) + int(math.Abs(float64(jumps[index-2])-float64(jumps[index])))
		}
	}

	minimum := math.Min(float64(lompat1), float64(lompat2))
	return int(minimum)
}
