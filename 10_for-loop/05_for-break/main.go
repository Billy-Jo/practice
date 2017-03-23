package main

import (
	"fmt"
)

func main() {
	i := 0
	for {
		fmt.Println(i)
		if i >= 10 {
			fmt.Println(i, "is over 10")
			break
		}

		i++
	}
}
