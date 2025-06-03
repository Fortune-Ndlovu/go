package main
import "fmt"

func main () {
	// var EvenNum[5] int
	// EvenNum[0] = 3
	// EvenNum[1] = 6
	// EvenNum[2] = 9
	// EvenNum[3] = 12
	// EvenNum[4] = 14
	// fmt.Println(EvenNum[2])
	
	// EvenNum := [5]int{3, 6, 9, 12, 14}
	// fmt.Println(EvenNum[2])

	// iterate through the array
	// EvenNum := [5]int{3, 6, 9, 12, 14}
	// for _, value := range EvenNum {
	// 	fmt.Println(value)
	// }

	// iterate through the array
	// EvenNum := [5]int{3, 6, 9, 12, 14}
	// for i, value := range EvenNum {
	// 	fmt.Println(value, i)
	// }

	// slice the array
	// EvenNum := [5]int{3, 6, 9, 12, 14}
	// sliced := EvenNum[3:5]
	// fmt.Println(sliced) // [12 14]

	numSlice := []int{5, 4, 3, 2, 1}

	sliced := numSlice[0:]
	fmt.Println(sliced) // [5 4 3 2 1]
	
	slice2 := make([]int, 5, 10)
	copy(slice2, numSlice)
	fmt.Println("slice2 ", slice2) // [5 4 3 2 1]

	slice3 := append(numSlice, 3, 0, -1)
	fmt.Println("slice3 ", slice3) // [5 4 3 2 1 3 0 -1]
}