package main

import "fmt"

func main () {
	a,b := 5,10
	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b)
	fmt.Println("a / b = ", a/b)
	fmt.Println("a > b = ", a>b)
	fmt.Println("a >= b = ", a>=b)
	fmt.Println("a == b = ", a==b)
	fmt.Println("a != b = ", a!=b)
	fmt.Println("a < b && a > b = ", a < b && a > b)
	fmt.Println("a > b || b > a = ", a > b || b > a)
	fmt.Println("!(a > b) = ", !(a > b))

	/* 
	In these lines, we're using **logical operators** to combine or invert **comparison results** between two integer variables `a` and `b`. The first line checks whether `a` is less than `b` **and** greater than `b` at the same time using the `&&` (AND) operator — which is always false because both conditions can’t be true together. The second line checks whether `a` is greater than `b` **or** `b` is greater than `a` using the `||` (OR) operator — this returns true if **either** condition is true. The third line uses the `!` (NOT) operator to **invert** the result of the comparison `a > b`; so if `a` is not greater than `b`, the expression evaluates to true. These lines demonstrate how Go allows combining multiple boolean expressions to form more complex logic decisions.
	*/
}