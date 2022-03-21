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

	// 무한 루프
	j := 0
	for {
		fmt.Print(" ", j) // 0 1 2 3 ... 100
		if j == 100 {
			break
		}
		j++
	}
	fmt.Println()

	// continue
	for k := 1; k < 100; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Print(" ", k) // 1 3 5 7 9 ... 99
	}
	fmt.Println()
}
