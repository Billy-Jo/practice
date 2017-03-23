package main

import "fmt"

func main() {
	greet("Jane", "Doe")
}

func greet(fname string, iname string) {
	fmt.Println(fname, iname)
}
