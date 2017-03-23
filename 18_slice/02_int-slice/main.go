package main

import (
	"fmt"
)

func main() {
	xs := []int{1, 3, 5, 7, 9, 11}
	for i, v := range xs {
		fmt.Println(i, " - ", v)
	}
}
