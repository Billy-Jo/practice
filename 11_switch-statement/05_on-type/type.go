package main

import "fmt"

//Contact dddd
type Contact struct {
	greeting string
	name     string
}

//SwitchOnType ddddd
func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	default:
		fmt.Println("unKnown")
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("McLeod")
	var t = Contact{"Good to see you", "Tim"}

	SwitchOnType(t)
	SwitchOnType(t.greeting)
	SwitchOnType(t.name)
	fmt.Println(1)
	fmt.Println(t)
	fmt.Println(t.greeting)
	fmt.Printf("%T \n", t)
}
