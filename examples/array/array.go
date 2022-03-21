package main

import (
	"fmt"
)

func main() {
	// 길이는 5로 고정
	var a [5]int
	for i := 0; i < 5; i++ {
		a[i] = i
	}
	fmt.Println(a) // [0 1 2 3 4]

	var b = [5]int{1, 2, 3, 4, 5}
	fmt.Println(b) // [1 2 3 4 5]

	c := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(c) // [a b c d e]

}
