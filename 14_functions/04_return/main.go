package main

import (
	"fmt"
)

func main() {
	fmt.Println(greet(11, 11))
}

func greet(fname, lname int) string {
	return fmt.Sprint(fname, lname)
}
