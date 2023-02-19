package main

import "fmt"

func main() {
	fmt.Println("apakah palindrome?")

	var kata, temp, temp2 string

	fmt.Print("masukkan kata: ")
	fmt.Scanf("%s", &kata)

	for i := 0; i < len(kata); i++ {
		temp += string(kata[i])
	}

	for j := len(kata) - 1; j >= 0; j-- {
		temp2 += string(kata[j])
	}

	fmt.Println("captured:", temp, temp2)

	if temp == temp2 {
		fmt.Println("palindrome")
	} else {
		fmt.Println("bukan palindrome")
	}
}
