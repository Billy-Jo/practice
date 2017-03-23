package main

import (
	"fmt"
)

func main() {
	switch "Jenny" {
	case "tim", "Jenny":
		fmt.Println("wassup tim or err Jenny")
	case "marcus", "medhi":
		fmt.Println("both of your names star twith m")
	case "Julian", "Sushant":
		fmt.Println("wassup Julian / student")
	}
}
