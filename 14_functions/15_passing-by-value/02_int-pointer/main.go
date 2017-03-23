package main

import "fmt"

func main() {

	age := 44
	fmt.Println(&age)
	fmt.Printf("%T\n", &age)
	changeMe(&age)

	fmt.Println(&age)
	fmt.Println(age)
}

func changeMe(z *int) {
	fmt.Println(z)
	fmt.Println(*z)
	*z = 24
	fmt.Println(z)
	fmt.Println(*z)
}
