package main

import "fmt"

func main() {
	fmt.Println(SegitigaPascal(1))
	fmt.Println(SegitigaPascal(2))
	fmt.Println(SegitigaPascal(3))
	fmt.Println(SegitigaPascal(4))
	fmt.Println(SegitigaPascal(5))
}

func SegitigaPascal(numRows int) [][]int {
	rows := make([][]int, numRows, numRows)

	for i := 0; i < len(rows); i++ {
		rows[i] = make([]int, i+1)
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || i == j {
				rows[i][j] = 1
			} else {
				rows[i][j] = rows[i-1][j-1] + rows[i-1][j-0]
			}
		}
	}

	return rows
}
