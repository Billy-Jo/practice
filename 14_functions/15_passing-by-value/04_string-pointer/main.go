package main

import (
	"fmt"
)

func main() {
	name := "todd"
	fmt.Println(&name)

	changeMe(&name)

	fmt.Println(&name)
	fmt.Println(name)
}

func changeMe(z *string) {
	fmt.Println(z)
	fmt.Println(*z)
	*z = "Rocky"
	fmt.Println(z)
	fmt.Println(*z)
	fmt.Println(2)
	fmt.Println(&z)
}
