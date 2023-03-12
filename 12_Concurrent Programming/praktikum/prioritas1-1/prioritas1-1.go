package main

import (
	"fmt"
	"time"
)

func main() {
	go angkaKelipatan(3, 5)
	time.Sleep(3 * time.Second)
}

func angkaKelipatan(x int, n int) {
	var arrKelipatan []int
	kelipatan := x

	for n > 0 {
		arrKelipatan = append(arrKelipatan, kelipatan)
		kelipatan += x
		n--
	}

	fmt.Println(arrKelipatan)
}
