package main

import (
	"fmt"
)

func main() {
	// range는 array, slice, map 등을 순회
	a := []string{"a", "b", "c", "d", "e"}
	for idx, val := range a {
		fmt.Println(idx, val) // 0 a, 1 b, 2 c, 3 d, 4 e
	}

	b := map[string]string{
		"apple": "fruit",
		"onion": "vegetable",
	}
	for key, val := range b {
		fmt.Println(key, val) // apple fruit, onion vegetable
	}

	for _, val := range b {
		fmt.Println(val) // fruit, vegetable
	}

	for key := range b {
		fmt.Println(key) // apple, onion
	}
}
