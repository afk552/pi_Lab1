package main

import "fmt"

func main() {
	fmt.Println("hello, world")
	fmt.Print("hello, world")
	// вывод будет в две строки:
	// hello, world
	// hello, world

	fmt.Print("Ivan", 27)   // Ivan27
	fmt.Println("Ivan", 27) // Ivan 27

	name := "Ivan"
	age := 27
	fmt.Println("My name is", name, "and I am", age, "years old.")
}
