package main
import "fmt"

func main () {
	fmt.Println(add(5,7))

// func add (num1, num int) int {
// 	return num1 + num

	// recursion calling a fun inside itself
	num := 5
	fmt.Println(factorial(num))
}
func add(x, y) int {
	return x + y
}

func factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}
