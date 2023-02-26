package main

import "fmt"

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(Mapping([]string{}))
}

func Mapping(slice []string) map[string]int {
	mapping := make(map[string]int)

	for _, value := range slice {
		var value2, keySudahAda = mapping[value]

		if keySudahAda {
			mapping[value] = value2 + 1
		} else {
			mapping[value] = value2 + 1
		}
	}

	return mapping
}
