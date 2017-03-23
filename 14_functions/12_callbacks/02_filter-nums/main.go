package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	fmt.Println("------")
	for _, n := range numbers {
		if callback(n) {
			fmt.Println("!")
			xs = append(xs, n)
		}
		fmt.Println("?")
	}
	fmt.Println("return")
	return xs
}

func main() {
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		fmt.Println("n : ", n)
		return n > 1
	})

	fmt.Println(xs)
}
