package main
import "fmt"

func main () {
	age := 18

	switch age {
	case 16: fmt.Println("Prep for college")
	case 18: fmt.Println("Find your purpose")
	case 23: fmt.Println("Progress in your job")
	default: fmt.Println("reset your life")
	}
}