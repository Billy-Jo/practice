package main

import "fmt"

func hello() {
	fmt.Println("hello")
}

func world() {
	fmt.Println("world")
}

func main() {
	world()
	hello()
	fmt.Println(world)
	fmt.Println(hello)
}
