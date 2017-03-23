package main

import (
	"fmt"
)

func main() {
	myGreeting := map[string]string{
		"Tim":    "good morning",
		"Janney": "Bonjour",
	}

	myGreeting["Harleen"] = "howdy"

	fmt.Println(myGreeting)
}
