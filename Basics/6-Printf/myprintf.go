package main

import "fmt"

func myprintft () {
	const pi float64 = 3.16412732
	x:=5
	isbool := true

	fmt.Printf(" %f \n ", pi)
	fmt.Printf(" %T \n ", isbool)
	fmt.Printf(" %t \n ", pi)
	fmt.Printf(" %d \n ", x)
}