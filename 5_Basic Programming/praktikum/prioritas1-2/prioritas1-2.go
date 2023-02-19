package main

import "fmt"

func main() {
	cekAngka := 21

	if cekAngka%2 == 0 {
		fmt.Println(cekAngka, "adalah angka genap")
	} else {
		fmt.Println(cekAngka, "adalah angka ganjil")
	}
}
