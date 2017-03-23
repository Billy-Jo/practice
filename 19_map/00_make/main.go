package main

import (
	"fmt"
)

func main() {
	var myGreeting = make(map[string]string)
	// var myMap = map[string]string{}
	myGreeting["Tim"] = "good morning"
	myGreeting["Jenny"] = "Bonjour"
	fmt.Println(len(myGreeting))
	//fmt.Println(cap(myGreeting))
	fmt.Printf("%T\n", myGreeting)
	fmt.Println(myGreeting)
}
