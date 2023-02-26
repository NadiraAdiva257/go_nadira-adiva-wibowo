package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(MunculSekali("76523752"))
	fmt.Println(MunculSekali("1122"))
	fmt.Println(MunculSekali("1234123"))
	fmt.Println(MunculSekali("12345"))
	fmt.Println(MunculSekali("1122334455"))
	fmt.Println(MunculSekali("0872504"))
}

func MunculSekali(angka string) []int {
	var arrayAngka []int
	mapping := make(map[string]int)

	for _, value := range angka {
		var value2, angkaSudahAda = mapping[string(value)]

		if angkaSudahAda {
			mapping[string(value)] = value2 + 1
		} else {
			mapping[string(value)] = value2 + 1
		}
	}

	for key, _ := range mapping {
		if mapping[key] == 1 {
			konvAngka, _ := strconv.Atoi(key)
			arrayAngka = append(arrayAngka, konvAngka)
		}
	}

	return arrayAngka
}
