package main

import (
	"fmt"
)

func main() {
	a := 43
	c := &a
	*c = 20

	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(*c)

	var b *int = &a

	fmt.Println(b)
}
