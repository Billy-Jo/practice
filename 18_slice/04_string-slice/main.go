package main

import (
	"fmt"
)

func main() {
	greeting := []string{
		"good morning",
		"bonjour",
		"dias",
		"oehdod",
		"selamat pagi!",
	}

	for i, currentEntry := range greeting {
		fmt.Println(i, currentEntry)
	}

	for i := 0; i < len(greeting); i++ {
		fmt.Println(i, greeting[i])
	}
}
