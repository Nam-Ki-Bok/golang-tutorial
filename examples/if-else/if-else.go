package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 일반적인 if-else
	if rand.Int()%2 == 0 {
		fmt.Println("짝수입니다.")
	} else {
		fmt.Println("홀수입니다.")
	}

	// 변수 선언과 동시에 사용 가능
	// 단, 변수는 if 안에서만 사용
	if num := rand.Int(); num%2 == 0 {
		fmt.Println("짝수입니다.")
	} else {
		fmt.Println("홀수입니다.")
	}
	// fmt.Println(num) -> error!
}
