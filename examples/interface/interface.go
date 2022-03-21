package main

import (
	"fmt"
)

type Person struct {
	name string
}

func main() {
	printAllType(1, "hello", 3.14, Person{name: "wade"})
}

// 빈 인터페이스를 통해 모든 타입을 대응
func printAllType(a ...interface{}) {
	for _, val := range a {
		switch val.(type) {
		case int:
			fmt.Println(val, "This type is int") // 1 This type is int
		case string:
			fmt.Println(val, " This type is string") // hello This type is string
		case float64:
			fmt.Println(val, " This type is float64") // 3.14 This type is float64
		case Person:
			fmt.Println(val, " This type is Person") // {wade} This type is Person
		}
	}
}
