package main

import (
	"fmt"
)

type Dog struct {
	name string
}

type Cat struct {
	name string
}

type animal interface {
	whoAmI() string
	bark()
}

func action(a animal) {
	a.bark()
	fmt.Println("hello", a.whoAmI())
}

func main() {
	dog := &Dog{name: "tory"}
	cat := &Cat{name: "nana"}

	action(dog)
	action(cat)
}

func (d *Dog) whoAmI() string {
	fmt.Println("I'm", d.name)
	return d.name
}

func (d *Dog) bark() {
	fmt.Println("mung!")
}

func (c *Cat) whoAmI() string {
	fmt.Println("I'm", c.name)
	return c.name
}

func (c *Cat) bark() {
	fmt.Println("meow!")
}
