package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

/*
Note: go run compiles your code into a tmp dir and run the binary build in that tmp dir and deletes the tmp dir after the program exits.
go run hello.go # Hello, World!

Note: go build build the binary in the current dir for later use
go build hello.go # hello
./hello # Hello, World!

Note: You can also build the binary with a different name
go build -o myhello hello.go # myhello
./myhello # Hello, World!
*/
