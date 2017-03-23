package main

import (
	"fmt"
)

func main() {
	myGreeting := map[int]string{
		0: "good morning",
		1: "Bonjour",
		2: "Buenos dias",
		3: "Bonginorano",
	}

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}
