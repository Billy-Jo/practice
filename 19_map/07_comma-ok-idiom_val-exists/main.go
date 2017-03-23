package main

import (
	"fmt"
)

func main() {
	myGreeting := map[int]string{
		0: "good morning",
		1: "bonjour",
		2: "buenos dias",
		3: "bonsssss",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, 2)

	if val, exist := myGreeting[2]; exist {

		fmt.Println("val : ", val)
		fmt.Println("exist: ", exist)
	} else {
		fmt.Println("that value doesn't exist")
		fmt.Printf("%T\n", val)
		fmt.Println("val : ", val)
		fmt.Println("exist : ", exist)
	}

	fmt.Println(myGreeting)
}
