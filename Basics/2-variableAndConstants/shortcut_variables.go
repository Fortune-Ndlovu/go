package main

import "fmt"

func shortcutVariables () {
	age := 23
	fmt.Println("My current age is: ", age)
}

// `:=` automatically figures out the type based on the value
// used inside functions only, not at the package level