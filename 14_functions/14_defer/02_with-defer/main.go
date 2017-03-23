package main

import "fmt"

func hello() {
	fmt.Println("hello")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world()
	defer fmt.Println("a")
	hello()
	world()
	defer hello()
	world()
	defer world()
	defer world()
}
