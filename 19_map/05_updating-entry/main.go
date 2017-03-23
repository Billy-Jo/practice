package main

import (
	"fmt"
)

func main() {
	myGreeting := map[string]string{
		"Tim":   "good morning",
		"Jenny": "Bounjour",
	}

	myGreeting["Harllen"] = "howdy"
	fmt.Println(myGreeting)
	myGreeting["Harleen"] = "gidday"
	fmt.Println(myGreeting)
}
