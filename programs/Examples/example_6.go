package main

import "fmt"

func main() {
	// basic operations
	a := 100
	b := 10
	c := a + b // с = 110
	fmt.Println(c)
	c = a * b // с = 1000
	fmt.Println(c)
	c = a - b // с = 90
	fmt.Println(c)
	c = a / b // с = 10
	fmt.Println(c)

	// division
	var m float32 = 10.0 / 6
	fmt.Println(m)

	// div
	var test int = 10 / 6
	fmt.Println(test)

	// mod
	var z int = 10 % 6
	fmt.Println(z)

	var i int = 1
	i++ //inc
	fmt.Println(i)
	i-- //dec
	fmt.Println(i)
}
