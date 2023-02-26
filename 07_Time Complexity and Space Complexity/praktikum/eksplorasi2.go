package main

import "fmt"

func main() {
	fmt.Println(pow(2, 3))
	fmt.Println(pow(5, 3))
	fmt.Println(pow(10, 2))
	fmt.Println(pow(2, 5))
	fmt.Println(pow(7, 3))
}

// maaf kak, belum menemukan cara supaya jadi O(log n)
func pow(x, n int) int {
	hasil := 1
	i := 0

	for n > 0 {
		hasil *= x
		i++

		if i == n {
			break
		}
	}

	return hasil
}
