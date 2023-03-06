package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(RepresentasiBiner(2)) // 0, 1, 10
	fmt.Println(RepresentasiBiner(3)) // 0, 1, 10, 11
	fmt.Println(RepresentasiBiner(7)) // 0, 1, 10, 11, 100, 101, 110, 111
	fmt.Println(RepresentasiBiner(9)) // 0, 1, 10, 11, 100, 101, 110, 111, 1000, 1001
}

func RepresentasiBiner(n int) []int {
	var ansKebalik []int
	var ans []int

	for n >= 0 {
		ansKebalik = append(ansKebalik, ConvertToBiner(n))
		n--
	}

	for i := len(ansKebalik) - 1; i >= 0; i-- {
		ans = append(ans, ansKebalik[i])
	}

	return ans
}

func ConvertToBiner(angka int) int {
	var binerKebalik string
	var biner string

	for angka > 0 {
		mod := angka % 2
		binerKebalik += strconv.Itoa(mod)
		angka /= 2
	}

	for i := len(binerKebalik) - 1; i >= 0; i-- {
		biner += string(binerKebalik[i])
	}

	binerInt, _ := strconv.Atoi(biner)
	return binerInt
}
