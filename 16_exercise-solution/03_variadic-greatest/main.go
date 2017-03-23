package main

import (
	"fmt"
)

func main() {
	greatest := max(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(greatest)
}

/*
func findGreatest(num ...interface{}) {
	var theGreatest int64
	for _, number := range num {
		fmt.Println(number)

	}
}
*/

func max(numbers ...int) int {
	fmt.Println(numbers)
	fmt.Printf("%T\n", numbers)
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}

	return largest
}
