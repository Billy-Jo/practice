package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"monday", "tuesday"}
	myOtherSlice := []string{"wednesday", "thursday", "friday"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)

	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println(mySlice)
}
