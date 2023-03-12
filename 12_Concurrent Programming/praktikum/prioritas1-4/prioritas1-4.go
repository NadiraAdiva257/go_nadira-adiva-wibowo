package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutex sync.Mutex
	ch := make(chan int)

	input := 5
	hasil := 1

	go func() {
		for i := input; i > 0; i-- {
			mutex.Lock()
			hasil *= i
			mutex.Unlock()
		}
		ch <- hasil
		close(ch)
	}()

	for value := range ch {
		fmt.Println(value)
	}
}
