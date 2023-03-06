package main

import "fmt"

func main() {
	fmt.Println(Fibo(0))  // 0
	fmt.Println(Fibo(1))  // 1
	fmt.Println(Fibo(2))  // 1
	fmt.Println(Fibo(3))  // 2
	fmt.Println(Fibo(5))  // 5
	fmt.Println(Fibo(6))  // 8
	fmt.Println(Fibo(7))  // 13
	fmt.Println(Fibo(9))  // 34
	fmt.Println(Fibo(10)) // 55
}

func Fibo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fibo(n-1) + Fibo(n-2)
	}
}
