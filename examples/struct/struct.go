package main

import (
	"fmt"
)

// Person
// 소문자로 시작하는 경우 외부에서 접근 불가
// 대문자로 시작하는 경우 외부에서 접근 가능
type Person struct {
	name string
	age  int
	etc  map[string]string
}

func newPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
		etc:  map[string]string{},
	}
}

func setName(name string, p *Person) {
	p.name = name
}

func (p *Person) Name() string {
	return p.name
}

func main() {
	p := newPerson("wade", 27)

	fmt.Println(p) // &{Wade 27 map[]}

	fmt.Println(p.Name()) // wade

	setName("foo", p)
	fmt.Println(p.Name()) // foo

}
