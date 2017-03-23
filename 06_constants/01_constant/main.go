package main

import (
	"fmt"
)

const p string = "death & taxes"

func main() {
	const q = 42

	fmt.Println("q - ", q)
	fmt.Println("p - ", p)
	foo(q)
}

func foo(e int) {
	fmt.Println(e)
}
