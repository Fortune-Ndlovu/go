package main
import "fmt"

func main () {
	// While loop
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// inner loop
	for i:=1; i<10; i++{
		for j:=1; j<i; j++ {
			fmt.Printf("*")
		}
		fmt.Println(i)
	}
}