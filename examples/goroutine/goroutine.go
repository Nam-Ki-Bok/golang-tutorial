package main

import (
	"fmt"
	"time"
)

func main() {
	// 1초마다 loop1, 2, 3이 동시에 실행
	// 보장 되지 않는 실행순서
	go loop1()
	go loop2()
	go loop3()

	time.Sleep(time.Second * 350)
}

func loop1() {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
		fmt.Println("I'm loop 1", i)
	}
}

func loop2() {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
		fmt.Println("I'm loop 2", i)
	}
}

func loop3() {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
		fmt.Println("I'm loop 3", i)
	}
}
