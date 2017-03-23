package main

import (
	"fmt"
)

func main() {
	greeting := make([]string, 3, 5)

	greeting[0] = "good morning"
	greeting[1] = "bonjour"
	greeting[2] = "buenos dias"
	greeting = append(greeting, "superbadham")
	greeting = append(greeting, "zao an")
	greeting = append(greeting, "ohayou")
	greeting = append(greeting, "gidday")

	fmt.Println(greeting[0])
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))
}
