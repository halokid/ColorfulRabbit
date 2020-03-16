package main

import "fmt"

func Foo() {
	fmt.Println("Foo...")

	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
}

func main() {
	Foo()
}






