package main

import "fmt"

func main () {
	var name string = "Fortune"
	const pi float64 = 3.141234
	light := true
	x := 5

	fmt.Println(len(name))
	fmt.Println(name + " is a chill dude")

	fmt.Printf("%.4f \n", pi) // 3.1412

	fmt.Printf("%T \n", name) // type: string

	fmt.Printf("%t \n", light) // true
	fmt.Printf("%d \n", x) // 5 (value)
	fmt.Printf("%b \n", 25) // 11001 (binary)
	fmt.Printf("%c \n", 34) // "


}