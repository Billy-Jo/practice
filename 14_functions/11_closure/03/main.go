package main

import "fmt"

func main() {
	x := 0
	increament := func() int {
		x++
		return x
	}

	fmt.Println(increament())
	fmt.Println(increament())

	fmt.Printf("%T\n", increament)
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope

anonymous function
a function without a name

func expression
assigning a func to a variable
*/
