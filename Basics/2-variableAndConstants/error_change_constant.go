package main

import "fmt"

func errorChangeConstant () {
	const name = "Fortune"
	// name = "lion"
	fmt.Println("My name: ", name)
}