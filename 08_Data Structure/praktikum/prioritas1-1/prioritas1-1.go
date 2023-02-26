package main

import "fmt"

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}

func ArrayMerge(arrayA, arrayB []string) []string {
	var arrayGabungan2 []string

	arrayGabungan := append(arrayA, arrayB...)
	mapping := make(map[string]int)

	for _, value := range arrayGabungan {
		var _, namaSudahAda = mapping[value]

		if namaSudahAda {
			continue
		} else {
			mapping[value] = 1
			arrayGabungan2 = append(arrayGabungan2, value)
		}
	}

	return arrayGabungan2
}
