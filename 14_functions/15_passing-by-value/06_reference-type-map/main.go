package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	changeME(m)
	fmt.Println(m["todd"])
}

func changeME(z map[string]int) {
	z["todd"] = 44
}
