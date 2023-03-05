package main

import (
	"fmt"
)

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
	fmt.Println(primeX(15)) // 47
}

func primeX(number int) int {
	counter := 0
	result := 0

	for i := 2; counter < number; i++ {
		if cekPrime(i) {
			result = i
			counter++
		}
	}

	return result
}

func cekPrime(number int) bool {
	temp := 0
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			temp += 1
		}
	}

	if temp == 2 {
		return true
	} else {
		return false
	}
}
