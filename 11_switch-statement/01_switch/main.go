package main

import (
	"fmt"
)

func main() {
	switch "Medhi" {
	case "daniel":
		fmt.Println("massup daniel")
	case "Medhi":
		fmt.Println("messup medhi")
	case "jenny":
		fmt.Println("messup jenny")
	default:
		fmt.Println("have you no friends?")
	}
}
