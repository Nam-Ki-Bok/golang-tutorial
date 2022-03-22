package main

import (
	"fmt"
)

// int 채널
var ch = make(chan int)

func main() {
	a, b := 10, 5
	go func() {
		// 수신을 할 때까지 대기
		ch <- a + b
	}()

	// 송신이 올 때까지 대기
	fmt.Println(<-ch) // 15
}
