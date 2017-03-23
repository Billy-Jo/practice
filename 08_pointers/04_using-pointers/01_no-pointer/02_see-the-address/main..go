package main

import (
	"fmt"
)

func zero(z int) {
	fmt.Printf("inside func zero x = %p\n", &z)
	fmt.Println("inside func zero x =", &z)
	z = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x)
	fmt.Println(&x)
	zero(x)
	fmt.Println(x)

}
