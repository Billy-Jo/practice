package main

import (
	"fmt"
)

func max(x int) int {
	return 42 + x
}

func main() {
	var max = max(7)
	fmt.Println(max)

	fmt.Println(max(7))
}
