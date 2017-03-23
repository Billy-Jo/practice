package main

import (
	"fmt"
)

func main() {
	myGreeting := map[string]string{
		"tim":   "good morning",
		"jenny": "Bonjour",
	}

	fmt.Println(myGreeting["jenny"])
}
