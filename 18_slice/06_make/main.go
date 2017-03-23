package main

import "fmt"

func main() {
	customerNumber := make([]int, 3)
	// 3 is length & capacity
	// length - number of elements refered to by the slice
	// capacity - number of elements in the underlying array

	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 15

	fmt.Println(customerNumber[0])
	fmt.Println(customerNumber[1])
	fmt.Println(customerNumber[2])

	greetirng := make([]string, 3, 5)

	greetirng[0] = "good morning"
	greetirng[1] = "bonjour"
	greetirng[2] = "dias"

	fmt.Println(greetirng[2])
}
