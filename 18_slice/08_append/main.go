package main

import "fmt"

func main() {
	greetring := make([]string, 3, 5)

	greetring[0] = "good morning"
	greetring[1] = "bonjour"
	greetring[2] = "buenos dias"
	greetring = append(greetring, "suprabadham")

	fmt.Println(greetring[3])
}
