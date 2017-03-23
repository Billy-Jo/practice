package main

import (
	"fmt"
)

func main() {

	myFriendsName := "Ma"
	fmt.Println(len("%%%"))
	switch {
	case len(myFriendsName) == 2:
		fmt.Println("wassup my friend name of length 2")
	case myFriendsName == "Time":
		fmt.Println("wassup tim")
	case myFriendsName == "jenny":
		fmt.Println("wassup Jenny")
	case myFriendsName == "marcus", myFriendsName == "medhi":
		fmt.Println("marcus medhi")
	case myFriendsName == "julian":
		fmt.Println("wassup julian")
	default:
		fmt.Println("this is the default")
	}
}
