package main

import "fmt"

func makeGreeter() func() string {
	fmt.Println(5)
	return func() string {
		fmt.Println(1)
		return "hello world"
	}
}

func main() {
	fmt.Println(2)
	greet := makeGreeter()
	fmt.Println(3)
	fmt.Println(greet())
	fmt.Printf("%T\n", greet)
	fmt.Println(4)
}
