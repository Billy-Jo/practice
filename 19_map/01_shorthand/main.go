package main

import (
	"fmt"
)

func main() {
	myGreeting := make(map[string]string)
	myGreeting["TIm0"] = "good morning"
	myGreeting["Jenny"] = "BOnjour"
	fmt.Println(myGreeting)
}
