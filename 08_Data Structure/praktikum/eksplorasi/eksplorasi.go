package main

import (
	"fmt"
	"math"
)

func main() {
	var array = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	diagonalKiriKeKanan := 0
	diagonalKananKeKiri := 0

	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if i == j {
				diagonalKiriKeKanan += array[i][j]
			}

			if j == (len(array)-1)-i {
				diagonalKananKeKiri += array[i][j]
			}
		}
	}

	fmt.Println("diagonalKiriKeKanan:", diagonalKiriKeKanan)
	fmt.Println("diagonalKananKeKiri:", diagonalKananKeKiri)

	selisihAbsolutDiagonal := math.Abs(float64(diagonalKiriKeKanan) - float64(diagonalKananKeKiri))
	fmt.Println("selisih absolut diagonal:", selisihAbsolutDiagonal)
}
