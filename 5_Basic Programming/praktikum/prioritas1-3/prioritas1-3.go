package main

import "fmt"

func main() {
	nilai := 75

	if nilai >= 80 && nilai <= 100 {
		fmt.Println("nilai kamu adalah A")
	} else if nilai >= 65 && nilai <= 79 {
		fmt.Println("nilai kamu adalah B")
	} else if nilai >= 50 && nilai <= 64 {
		fmt.Println("nilai kamu adalah C")
	} else if nilai >= 35 && nilai <= 49 {
		fmt.Println("nilai kamu adalah D")
	} else if nilai >= 0 && nilai <= 34 {
		fmt.Println("nilai kamu adalah E")
	} else {
		fmt.Println("nilai invalid")
	}
}
