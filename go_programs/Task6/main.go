/*
Задание 6. Задача: Дано натуральное число, выведите его последнюю цифру.
На вход дается натуральное число N, не превосходящее 10000.
Выведите одно целое число - ответ на задачу.
*/

package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	number %= 10 // number = number % 10
	fmt.Println(number)
}
