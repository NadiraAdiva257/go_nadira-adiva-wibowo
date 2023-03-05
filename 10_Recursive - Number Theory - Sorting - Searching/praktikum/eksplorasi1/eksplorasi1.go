package main

import "fmt"

func main() {
	MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) // 6
	MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6})   // 7
	MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3})   // 7
	MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6})   // 8
	MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6})    // 12
}

func MaxSequence(arr []int) {
	tempMax := 0
	totalMax := 0

	for _, value := range arr {
		tempMax += value
		if tempMax < 0 {
			tempMax = 0
		} else {
			if totalMax < tempMax {
				totalMax = tempMax
			}
		}
	}

	fmt.Println(totalMax)
}
