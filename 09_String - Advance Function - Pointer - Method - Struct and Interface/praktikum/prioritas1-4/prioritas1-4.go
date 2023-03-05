package main

import (
	"fmt"
)

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	min, max = GetMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("nilai minimum adalah", min)
	fmt.Println("nilai maximum adalah", max)
}

func GetMinMax(number ...*int) (min, max int) {
	min = *number[0]
	max = *number[0]

	for _, value := range number {
		if *value < min {
			min = *value
		}

		if *value > max {
			max = *value
		}
	}

	return
}
