package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = []int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]

	a = append(a, 6) // 원소 추가
	fmt.Println(a)   // [1 2 3 4 5 6]

	var tmp = []int{7, 8, 9, 10}
	a = append(a, tmp...) // 슬라이스 추가
	fmt.Println(a)        // [1 2 3 4 5 6 7 8 9 10]

	b := make([]string, 5) // 길이가 5인 빈 문자열을 원소로 가지고 있는 slice 생성
	fmt.Println(b)         // [   ]

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

	c[0] = "A"     // c 값을 변경해도 b 값은 변하지 않는다
	fmt.Println(b) // [a b c d e]
	fmt.Println(c) // [A b c]

	d := []int{5, 4, 2, 7, 1, 9, 3, 6, 8}
	fmt.Println(sort.IntsAreSorted(d)) // false
	sort.Ints(d)                       // sort slice
	fmt.Println(sort.IntsAreSorted(d)) // true
}
