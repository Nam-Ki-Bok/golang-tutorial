package main

import (
	"fmt"
)

func main() {
	// 일반적인 switch
	// 해당 case 실행 뒤 종료
	switch fruit := ""; fruit {
	case "apple":
		fmt.Println("I am apple!")
	case "banana":
		fmt.Println("I am banana!")
	case "carrot", "onion":
		fmt.Println("I am vegetable!")
	default:
		fmt.Println("Who am I?")
	}

	// fallthrough 사용 시
	// 해당 case 밑에 있는 모든 case 실행
	switch num := 1; num {
	case 1:
		fmt.Println("1 continue..")
		fallthrough
	case 2:
		fmt.Println("2 continue..")
		fallthrough
	case 3:
		fmt.Println("3 finish!")
	}
}
