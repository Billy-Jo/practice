package main

import (
	"fmt"
)

func main() {
	myGreeting := map[string]string{
		"Tim":   "good morning",
		"Jenny": "Bonjour",
	}

	myGreeting["Harleen"] = "howdy"
	fmt.Println(len(myGreeting))

}
