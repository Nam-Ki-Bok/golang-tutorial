package main

import (
	"fmt"
)

// 참조에 의한 호출
func changeNamePtr(name *string) {
	*name = "I am pointer!"
}

// 값에 의한 호출
func changeNameVal(name string) {
	name = "I am value!"
}

func main() {
	name := "wade"

	changeNameVal(name)
	fmt.Println(name) // wade

	changeNamePtr(&name)
	fmt.Println(name) // I am pointer!

	fmt.Println(&name) // name이 저장 된 메모리 주소
}
