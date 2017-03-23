package main

import (
	"fmt"
)

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)
	student[0] = "todd"
	student = append(student, "ffs")
	fmt.Println(len(student))
	fmt.Println(cap(student))
	student = append(student, "ffs")
	student = append(student, "ffs")
	student = append(student, "ffs")
	student = append(student, "ffs")
	student = append(student, "ffs")
	student = append(student, "ffs")
	student = append(student, "ffs")
	student = append(student, "ffs")
	student = append(student, "ffs")
	student = append(student, "ffs")
	fmt.Println(len(student))
	fmt.Println(cap(student))
	student[1] = "billy"
	student = append(student, "ffs")
	student = append(student, "ffs")
	student = append(student, "ffs")
	student = append(student, "ffs")
	student = append(student, "ffs")
	student = append(student, "ffs")
	student = append(student, "ffs")
	student = append(student, "ffs")
	student = append(student, "ffs")
	student = append(student, "ffs")
	student = append(student, "ffs")
	fmt.Println(student)
	fmt.Println(student[0])
	fmt.Println(student[1])
	fmt.Println(len(student))
	fmt.Println(cap(student))
	fmt.Println("==========================")
	fmt.Println(len(student))
	fmt.Println(len(student[0]))
	fmt.Println(len(student[1]))
	fmt.Println("==========================")
	for i, v := range student {
		fmt.Println(i, v)
	}
	fmt.Println("==========================")
	fmt.Println(students)
	fmt.Println(student == nil)
}
