package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 3; i < 50; i += 3 {
			ch <- i
		}
		close(ch)
	}()

	for value := range ch {
		fmt.Print(value, " ")
	}
}
