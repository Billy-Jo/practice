package main

import (
	"fmt"
)

func main() {
	var result []int
	fmt.Println(result)

	mySlice := []string{"a", "b", "c", "d", "r", "h"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])
	fmt.Println(mySlice[2])
	fmt.Println("mySlice"[2])
}
