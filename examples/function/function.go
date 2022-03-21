package main

import (
	"errors"
)

var (
	ErrDivisionZero = errors.New("division by zero")
)

func plus(a, b int) int {
	return a + b
}

func multi(a, b int) (result int) {
	result = a * b
	return
}

func division(a, b int) (int, error) {
	if b == 0 {
		return -1, ErrDivisionZero
	}

	result := a / b
	return result, nil
}

func main() {
}
