package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("유저이름을 입력해주세요: ")
	result, err := fmt.Scan(&name)
	fmt.Println("hello", name)
	fmt.Println(result, err)
}
