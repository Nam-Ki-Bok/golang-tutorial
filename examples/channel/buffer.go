package main

import (
	"fmt"
	"time"
)

// int 버퍼 채널
var queue = make(chan int, 10)

// bool 채널
var done = make(chan bool)

func main() {
	// 수신을 기다리지 않고 채널에 모든 값을 삽입 후 다른 작업 수행
	go push()

	// 채널에 있는 값을 1초마다 가져온다
	go pop()

	// 송신이 올 때까지 대기
	<-done
}

func push() {
	for i := 0; i < 10; i++ {
		queue <- i
	}
	fmt.Println("push finished and doing another job")
}

func pop() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println(<-queue)
	}
	done <- true
}
