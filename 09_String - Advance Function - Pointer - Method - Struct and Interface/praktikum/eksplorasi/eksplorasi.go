package main

import "fmt"

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student's name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of student's name " + a.name + " is : " + c.Decode())
	}
}

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""
	const hurufUrut string = "abcdefghijklmnopqrstuvwxyz"
	const monoalphabetic string = "pdjnwquegyxklaczbmrsftoivh"
	var tempIndexHuruf []int

	for _, value := range s.name {
		for key2, value2 := range monoalphabetic {
			if string(value) == string(value2) {
				tempIndexHuruf = append(tempIndexHuruf, key2)
			}
		}
	}

	for _, value := range tempIndexHuruf {
		nameEncode += string(hurufUrut[value])
	}

	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""
	const hurufUrut string = "abcdefghijklmnopqrstuvwxyz"
	const monoalphabetic string = "pdjnwquegyxklaczbmrsftoivh"
	var tempIndexHuruf []int

	for _, value := range s.name {
		for key2, value2 := range hurufUrut {
			if string(value) == string(value2) {
				tempIndexHuruf = append(tempIndexHuruf, key2)
			}
		}
	}

	for _, value := range tempIndexHuruf {
		nameDecode += string(monoalphabetic[value])
	}

	return nameDecode
}
