package main

import (
	"fmt"
	"time"
)

func main() {
	go KemunculanTiapHuruf("lorem ipsum")
	go KemunculanTiapHuruf("nadira adiva")

	time.Sleep(1 * time.Second)
}

func KemunculanTiapHuruf(kalimat string) {
	mapping := make(map[string]int)

	for _, value := range kalimat {
		var value2, keySudahAda = mapping[string(value)]

		if keySudahAda {
			mapping[string(value)] = value2 + 1
		} else {
			mapping[string(value)] = value2 + 1
		}
	}

	fmt.Println(mapping)
}
