package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutex sync.Mutex
	ch := make(chan map[string]int)

	mapping := make(map[string]int)
	inputKalimat := "nadira adiva"

	go func(kalimat string) {
		for _, value := range kalimat {
			var value2, keySudahAda = mapping[string(value)]

			mutex.Lock()
			if keySudahAda {
				mapping[string(value)] = value2 + 1
			} else {
				mapping[string(value)] = value2 + 1
			}
			mutex.Unlock()
		}
		ch <- mapping
		close(ch)
	}(inputKalimat)

	for value := range ch {
		fmt.Println(value)
	}
}
