package main

import (
	"errors"
	"fmt"
)

var (
	ErrDivisionZero = errors.New("division by zero")
)

func main() {
	defer func() {
		// 프로그램 복구
		if r := recover(); r != nil {
			fmt.Println(r)
		}
		// panic 발생 후 main 함수 재호출
		main()
	}()

	result, err := division(10, 0)
	if err != nil {
		// 에러 발생으로 인한 프로그램 종료
		panic(err)
	}
	fmt.Println(result)
}

func division(a, b int) (int, error) {
	if b == 0 {
		return -1, ErrDivisionZero
	}

	return a / b, nil
}
