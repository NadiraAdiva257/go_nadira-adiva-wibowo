package main

import "fmt"

func main() {
	const liter = 1.5
	car1 := Car{
		tipeBahanBakar:  "solar",
		literBahanBakar: liter,
	}
	fmt.Println("perkiraan jarak tempuh dengan tipe bahan bakar", car1.tipeBahanBakar,
		"sebanyak", car1.literBahanBakar, "liter adalah", car1.Jarak())

	car2 := Car{
		tipeBahanBakar:  "pertamax",
		literBahanBakar: liter,
	}
	fmt.Println("perkiraan jarak tempuh dengan tipe bahan bakar", car2.tipeBahanBakar,
		"sebanyak", car2.literBahanBakar, "liter adalah", car2.Jarak())
}

type Car struct {
	tipeBahanBakar  string
	literBahanBakar float32
}

func (car Car) Jarak() int {
	var jarak int

	if car.tipeBahanBakar == "pertamax" {
		jarak = 200
	} else if car.tipeBahanBakar == "pertalite" {
		jarak = 150
	} else if car.tipeBahanBakar == "solar" {
		jarak = 100
	} else {
		jarak = 0
	}

	return jarak
}
