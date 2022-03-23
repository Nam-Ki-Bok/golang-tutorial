package main

import (
	"fmt"
)

type Adult struct {
	name string
}

type Baby struct {
	name string
}

type Person interface {
	eat() string
	think() string
}

func define(p Person) {
	fmt.Println("사람인..", p.eat())
	fmt.Println("사람인..", p.think())
}

func main() {
	a := new(Adult)
	b := new(Baby)

	// a, b 모두 Person interface의 메소드를 구현하였기 때문에
	// Person interface 타입으로 본다.
	define(a)
	define(b)
}

func (a *Adult) eat() string {
	return "어른이 밥을 먹습니다."
}

func (a *Adult) think() string {
	return "어른이 생각을 합니다."
}

func (b *Baby) eat() string {
	return "아기가 밥을 먹습니다."
}

func (b *Baby) think() string {
	return "아기가 생각을 합니다."
}
