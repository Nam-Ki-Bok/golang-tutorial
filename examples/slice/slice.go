package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]
	a = append(a, 6)
	fmt.Println(a) // [1 2 3 4 5 6]

	// length = 5
	b := make([]string, 5)
	fmt.Println(b) // [   ]
	b[0], b[1], b[2], b[3], b[4] = "a", "b", "c", "d", "e"
	fmt.Println(b) // [a b c d e]
}
