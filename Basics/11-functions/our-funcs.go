package main
import "fmt"

func main () {
// 	x,y := 5,6
// 	fmt.Println(add(x,y))
// }

// func add (num1, num int) int {
// 	return num1 + num

	// recursion calling a fun inside itself
	num := 5
	fmt.Println(factorial(num))
}

func factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}