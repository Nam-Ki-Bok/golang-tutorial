package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]

	// append
	// 원소 추가
	a = append(a, 6)
	fmt.Println(a) // [1 2 3 4 5 6]

	// 길이가 5인 빈 문자열을 원소로 가지고 있는 slice 생성
	b := make([]string, 5)
	fmt.Println(b) // [   ]

	b[0], b[1], b[2], b[3], b[4] = "a", "b", "c", "d", "e"
	fmt.Println(b) // [a b c d e]

	// slice[start:end]를 통해 원하는 범위 출력 가능
	// start index 부터 end - 1 index 까지 출력
	fmt.Println(b[:2])  // [a b]
	fmt.Println(b[2:])  // [c d e]
	fmt.Println(b[1:4]) // [b c d]

	c := make([]string, 3)
	copy(c, b)     // slice b를 slice c에 deep copy
	fmt.Println(c) // [a b c]

	// c 값을 변경해도 b 값은 변하지 않는다
	c[0] = "A"
	fmt.Println(b) // [a b c d e]
	fmt.Println(c) // [A b c]
}
