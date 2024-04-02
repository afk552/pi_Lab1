/*
Задание 8. Задача: Часовая стрелка повернулась с начала суток на d градусов.
Определите, сколько сейчас целых часов h и целых минут m.
На вход программе подается целое число d (0 < d < 360).
Выведите на экран фразу: It is ... hours ... minutes.
*/

package main

import "fmt"

func main() {
	var d int
	fmt.Scan(&d)

	hours := d * 2 / 60
	minutes := d * 2 % 60
	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}
