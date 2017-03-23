package main

import (
	"fmt"
)

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	var b int
	var c int

	fmt.Print("Enter meters swam: ")
	result, er := fmt.Scan(&meters, &b, &c)
	yards := meters * metersToYards
	fmt.Println(meters, "meters is ", yards, "yards.")
	fmt.Println("result -", result, " er -", er)
}
