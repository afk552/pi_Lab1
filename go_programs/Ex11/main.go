package main

import "fmt"

/*
const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)
*/

const (
	Sunday = iota // iota идентификатор Go для упрощения определений увеличивающихся чисел.
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Sunday)   // вывод 0
	fmt.Println(Saturday) // вывод 6
}
