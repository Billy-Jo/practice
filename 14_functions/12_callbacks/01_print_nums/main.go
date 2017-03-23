package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	for i, n := range numbers {
		fmt.Println("i: ", i, " n: ", n)
		callback(n)
	}
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
		fmt.Println("!")
	})

	for i := 1; i <= 10; i++ {
		fmt.Println("j : ", i)
	}
}
