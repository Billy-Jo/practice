package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	n := average(data...)
	fmt.Println(n)
}

func average(a ...float64) float64 {
	total := 0.0
	for c, v := range a {
		fmt.Println("c: ", c)
		fmt.Printf("c: %T \n", c)
		total += v
	}
	fmt.Println(a)
	fmt.Printf("%T \n", a)
	fmt.Println(float64(len(a)))
	fmt.Println(total)
	return total / float64(len(a))
}
