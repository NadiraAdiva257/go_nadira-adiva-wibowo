package main

import "fmt"

func main() {
	fmt.Println(KonversiRomawi(4))
	fmt.Println(KonversiRomawi(9))
	fmt.Println(KonversiRomawi(23))
	fmt.Println(KonversiRomawi(2021))
	fmt.Println(KonversiRomawi(1646))
	fmt.Println(KonversiRomawi(1555))
}

func KonversiRomawi(angka int) string {
	var banyak int
	var hasil string

	romawi := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}

	for angka > 0 {
		if angka >= 1000 {
			banyak = (angka - (angka % 1000)) / 1000
			for banyak > 0 {
				hasil += romawi[1000]
				banyak--
			}
			angka -= 1000
		}

		if angka >= 500 {
			if angka >= 900 {
				banyak = (angka - (angka % 900)) / 900
				for banyak > 0 {
					hasil += romawi[100] + romawi[1000]
					banyak--
				}
				angka -= 900
			} else {
				banyak = (angka - (angka % 500)) / 500
				for banyak > 0 {
					hasil += romawi[500]
					banyak--
				}
				angka -= 500
			}
		}

		if angka >= 100 {
			if angka >= 400 {
				banyak = (angka - (angka % 400)) / 400
				for banyak > 0 {
					hasil += romawi[100] + romawi[500]
					banyak--
				}
				angka -= 400
			} else {
				banyak = (angka - (angka % 100)) / 100
				for banyak > 0 {
					hasil += romawi[100]
					banyak--
				}
				angka -= 100
			}
		}

		if angka >= 50 {
			if angka >= 90 {
				banyak = (angka - (angka % 90)) / 90
				for banyak > 0 {
					hasil += romawi[10] + romawi[100]
					banyak--
				}
				angka -= 90
			} else {
				banyak = (angka - (angka % 50)) / 50
				for banyak > 0 {
					hasil += romawi[50]
					banyak--
				}
				angka -= 50
			}
		}

		if angka >= 10 {
			if angka >= 40 {
				banyak = (angka - (angka % 40)) / 40
				for banyak > 0 {
					hasil += romawi[10] + romawi[50]
					banyak--
				}
				angka -= 40
			} else {
				banyak = (angka - (angka % 10)) / 10
				for banyak > 0 {
					hasil += romawi[10]
					banyak--
				}
				angka -= 10
			}
		}

		if angka >= 5 {
			if angka >= 9 {
				banyak = (angka - (angka % 9)) / 9
				for banyak > 0 {
					hasil += romawi[1] + romawi[10]
					banyak--
				}
				angka -= 9
			} else {
				banyak = (angka - (angka % 5)) / 5
				for banyak > 0 {
					hasil += romawi[5]
					banyak--
				}
				angka -= 5
			}
		}

		if angka >= 1 {
			if angka >= 4 {
				banyak = (angka - (angka % 4)) / 4
				for banyak > 0 {
					hasil += romawi[1] + romawi[5]
					banyak--
				}
				angka -= 4
			} else {
				banyak = (angka - (angka % 1)) / 1
				for banyak > 0 {
					hasil += romawi[1]
					banyak--
				}
				angka -= 1
			}
		}
	}

	return hasil
}
