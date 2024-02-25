// Задание 5. По данному целому числу, найдите его квадрат.

package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)
	// fmt.Println(Math.Pow(number, 2))
	number *= number
	println(number)
}
