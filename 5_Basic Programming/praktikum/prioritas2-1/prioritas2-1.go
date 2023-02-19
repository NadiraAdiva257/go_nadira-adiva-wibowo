package main

import "fmt"

func main() {
	var input int

	fmt.Print("masukkan input: ")
	fmt.Scanln(&input)

	for i := 1; i <= input; i++ {
		for j := 1; j <= input-i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= i; k++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}
}
