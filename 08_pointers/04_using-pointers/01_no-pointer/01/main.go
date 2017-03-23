package main

import (
	"fmt"
)

func zero(x int) {
	x = 0
	fmt.Println("inside func zero x =", &x)
	fmt.Println("inside func zero x =", x)
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x)
	fmt.Println(&x)
}
