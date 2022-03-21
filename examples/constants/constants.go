package main

import (
	"fmt"
)

func main() {
	// 상수는 선언과 동시에 할당
	const a = 1
	const b = "banana"
	fmt.Println(a, b) // 1 banana

	// 선언 후 할당 시 에러 발생
	// const c string
	// c = "carrot" -> error!

	// enum
	// 0부터 1씩 증가하며 할당
	const (
		mon = iota
		tue
		wed
		thu
		fri
		sat
		sun
	)
	fmt.Println(mon, tue, wed, thu, fri, sat, sun) // 0, 1, 2, 3, 4, 5, 6
}
