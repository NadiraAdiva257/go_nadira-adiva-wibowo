package calculator

import "errors"

func addition(angka1, angka2 int) int {
	result := angka1 + angka2
	return result
}

func subtraction(angka1, angka2 int) int {
	result := angka1 - angka2
	return result
}

func multiplication(angka1, angka2 int) int {
	result := angka1 * angka2
	return result
}

func division(angka1, angka2 int) (int, error) {
	if angka2 == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	result := angka1 / angka2
	return result, nil
}
