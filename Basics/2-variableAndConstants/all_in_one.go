package main

import "fmt"

func allInOne () {
	var age int = 23
	const name string = "Fortune"
	var runSpeed float64 = 20.33

	// We can declare in one var
	var (
		a = 1
		b = 2
		c = 3
	)

	// popular way of delaring variables while assigning values
	x,y := 14,15

	fmt.Println("all in one")
	fmt.Println("Age", age)
	fmt.Println("Name", name)
	fmt.Println("run speed", runSpeed)
	fmt.Println("We were declared in a single var: ", a, b, c )
	fmt.Println("We were declared in one line the shorthand way: ", x, y )

}