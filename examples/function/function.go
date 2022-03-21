package main

import (
	"errors"
	"fmt"
)

var (
	ErrDivisionZero = errors.New("division by zero")
)

// 일반적인 함수 선언
func plus(a, b int) int {
	return a + b
}

// 반환 값을 미리 선언하는 함수
func multi(a, b int) (result int) {
	result = a * b
	return
}

// 다중 반환 함수
func division(a, b int) (int, error) {
	if b == 0 {
		return -1, ErrDivisionZero
	}

	result := a / b
	return result, nil
}

func main() {
	fmt.Println(plus(20, 10))     // 30
	fmt.Println(multi(20, 10))    // 200
	fmt.Println(division(20, 10)) // 2 nil
	fmt.Println(division(20, 0))  // -1 division by zero
}
