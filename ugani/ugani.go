package main

import (
	"fmt"
	"golang.org/x/example/hello/reverse"
)

func main() {
 // reversing both a string type and an Int type
	fmt.Println(reverse.String("Ugani Friend"), reverse.Int(24601))
}