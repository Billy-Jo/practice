package main

import "fmt"

func main() {
	var num1, num2 float32
	var result float32
	fmt.Println("작은 수 하나와 큰 수 하나를 입력해 주세요")
	input, err := fmt.Scan(&num1, &num2)
	if num1 > num2 {
		result = num1 / num2
	} else {
		result = num2 / num1
	}

	fmt.Println("입력하신 두 수에서 큰 수를 작은 수로 나눈 결과는 ", result, "입니다")
	fmt.Println(input, err)
}
