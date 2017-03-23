package main

import (
	"fmt"
)

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Printf("%T\n", increment)
	fmt.Printf("%T\n", increment())
	fmt.Printf("%T\n", wrapper)
	increment1 := wrapper
	fmt.Printf("%T\n", increment1)
	fmt.Println(1, increment1)
	fmt.Printf("%T\n", increment1())
	fmt.Println(2, increment1())
	fmt.Printf("%T\n", wrapper())

}
