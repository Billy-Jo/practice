package main

import (
	"fmt"
)

func main() {
	myGreeting := map[int]string{
		0: "good morning",
		1: "BOnjour",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, 2)
	fmt.Println(myGreeting)
}
