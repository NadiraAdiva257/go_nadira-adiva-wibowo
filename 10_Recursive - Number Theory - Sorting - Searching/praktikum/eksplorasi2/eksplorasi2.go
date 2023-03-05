package main

import "fmt"

func main() {
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})      // 3
	MaximumBuyProduct(50000, []int{15000, 10000, 12000, 5000, 3000}) // 5
	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}) // 4
	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})   // 4
	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})           // 1
	MaximumBuyProduct(0, []int{10000, 30000})                        // 0
}

func MaximumBuyProduct(money int, productPrice []int) {
	sortedProductPrice := BubbleSort(productPrice)
	var jumlah int

	for i := 0; i < len(sortedProductPrice); i++ {
		if money >= sortedProductPrice[i] {
			money -= sortedProductPrice[i]
			jumlah += 1
		}
	}

	fmt.Println(jumlah)
}

func BubbleSort(array []int) []int {
	cek := false

	for !cek {
		cek = true
		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				temp := array[i]
				array[i] = array[i+1]
				array[i+1] = temp
				cek = false
			}
		}
	}

	return array
}
