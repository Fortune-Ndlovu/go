package main

import "fmt"

func main () {
	var name string = "Fortune"
	fmt.Println("Name:", name)

	changeVariables()
	allInOne()
}

// Basic theory
/**
A variable in Go is a named place in memory where you can data. You can change the value of a variable later in the program
`var` - key word to declare a variable
`name` - the name of the variable
`string` - the type (sequence of characters)
"Fortune" - the value stored
*/