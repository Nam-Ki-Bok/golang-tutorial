package main

import (
	"fmt"
)

func main() {
	// 일반적인 for
	for i := 0; i < 10; i++ {
		fmt.Print(" ", i) // 0 1 2 3 4 5 6 7 8 9
	}
	fmt.Println()

	// continue, break
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		} else {
			fmt.Println(i)
		}

		if i == 9 {
			break
		}
	}

	// 조건문만 존재하는 for
	n := 1
	for n < 100 {
		fmt.Println(n) // 1 2 4 8 16 32 64
		n *= 2
	}

	// 무한 루프
	for {
		fmt.Println("Hello! this is infinity loop")
	}
}
