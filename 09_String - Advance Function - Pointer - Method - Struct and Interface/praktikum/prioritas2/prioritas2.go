package main

import "fmt"

func main() {
	fmt.Println(Caesar(3, "abc"))                           // def
	fmt.Println(Caesar(2, "alta"))                          // cnvc
	fmt.Println(Caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(Caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(Caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
	fmt.Println(Caesar(2000, "abcdefghijklmnopqrstuvwxyz")) // yzabcdefghijklmnopqrstuvwx
}

func Caesar(offset int, input string) string {
	var temp string
	modOffset := offset % 26

	for _, value := range input {
		selisihPosisi := 122 - value

		if int(selisihPosisi) < modOffset {
			asciiPetikSatu := 96
			kurangBerapaOffset := modOffset - int(selisihPosisi)
			temp += string(asciiPetikSatu + kurangBerapaOffset)
		} else {
			temp += string(int(value) + modOffset)
		}
	}

	return temp
}
