package main

import "fmt"

func main () {
	x := 10               // x is a variable with value 10
	changeValue(&x)       // &x gives the memory address of x
	fmt.Println(x)        // x has been changed to 7
}

func changeValue(x *int) {
	*x = 7                // *x means: go to the address and set the value to 7
}
