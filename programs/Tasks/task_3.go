/*
Задание 3. Напишите программу, которая последовательно делает следующие
операции с введённым числом: 1. Число умножается на 2; 2. Затем к числу
прибавляется 100. После этого должен быть вывод получившегося числа на
экран.
*/

package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	number *= 2   // number = number * 2
	number += 100 // number = number + 100
	fmt.Println(number)
}
