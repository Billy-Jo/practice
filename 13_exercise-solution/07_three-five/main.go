package main

import "fmt"

func main() {
	result := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}

	fmt.Println("the sum of natural numbers of the multiples of three or five below 100 is ", result)
}
