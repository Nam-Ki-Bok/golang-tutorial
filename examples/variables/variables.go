package main

import (
	"fmt"
)

func main() {
	// 변수 선언 후 할당
	var a int
	a = 1
	fmt.Println(a) // 1

	var b float32
	b = 3.14
	fmt.Println(b) // 3.14

	var c string
	c = "carrot"
	fmt.Println(c) // carrot

	var d bool
	d = true
	fmt.Println(d) // true

	// 여러 변수 선언 후 할당
	var e, f, g int
	e, f, g = 2, 3, 4
	fmt.Println(e, f, g) // 2, 3, 4

	// 선언과 동시에 할당
	// 타입 추론을 통해 자료형 생략 가능
	var h = 5
	var i = "ice"
	fmt.Println(h, i) // 5 ice

	// 짧은 선언
	// 함수 내에서만 사용 가능
	// 타입 추론을 통해 자료형 생략 가능
	j := 6
	k := "key"
	fmt.Println(j, k) // 6 key

	// 변수를 묶어서 선언 및 할당
	// 타입 추론을 통해 자료형 생략 가능
	var (
		l = 7
		m = "melon"
		n = false
	)
	fmt.Println(l, m, n) // 7, melon, false
}
