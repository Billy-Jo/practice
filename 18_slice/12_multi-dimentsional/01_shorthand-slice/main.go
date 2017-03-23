package main

import (
	"fmt"
)

func main() {
	student := []string{}
	students := [][]string{}
	fmt.Println(student)
	fmt.Println(len(student))
	fmt.Println("---------------")
	student = append(student, "1")
	fmt.Println(len(student))
	fmt.Println(student)
	fmt.Println(cap(student))
	fmt.Println("---------------")
	student = append(student, "1")
	fmt.Println(len(student))
	fmt.Println(student)
	fmt.Println(cap(student))
	fmt.Println("---------------")
	student = append(student, "1")
	fmt.Println(len(student))
	fmt.Println(student)
	fmt.Println(cap(student))
	fmt.Println("---------------")
	fmt.Println(students)
	fmt.Println(len(students))
	fmt.Println(student == nil)
}
