package main

import "fmt"

func changeVariables () {
	var age int = 23
	fmt.Println("Age before: ", age)

	age = 26
	fmt.Println("Age after: ", age)
}

// `int` is the type for integers (Whole numbers)
// We can change age later as its variable