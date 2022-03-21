package main

import (
	"fmt"
)

func main() {
	// key: string, value: string
	a := make(map[string]string)
	a["apple"] = "fruit"
	a["carrot"] = "vegetable"
	fmt.Println(a) // map[apple:fruit carrot:vegetable]

	value, isExist := a["apple"]
	fmt.Println(value, isExist) // fruit true

	delete(a, "apple")
	fmt.Println(a) // map[carrot:vegetable]

	// 값이 없더라도 에러가 발생하지 않는다
	value, isExist = a["apple"]
	fmt.Println(value, isExist) // "" false

}
