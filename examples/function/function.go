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

// 가변 함수
func sum(arr ...int) int {
	total := 0
	for _, val := range arr {
		total += val
	}
	return total
}

func main() {
	fmt.Println(plus(20, 10))     // 30
	fmt.Println(multi(20, 10))    // 200
	fmt.Println(division(20, 10)) // 2 nil
	fmt.Println(division(20, 0))  // -1 division by zero

	fmt.Println(sum(1, 2, 3))       // 6
	fmt.Println(sum(1, 2, 3, 4, 5)) // 15

	a := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(a...)) // 15
}
