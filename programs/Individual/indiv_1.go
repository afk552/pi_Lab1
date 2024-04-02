/*
Индивидуальное задание 1 (8). Объем и площадь цилиндрической банки:
Задайте переменные для радиуса и высоты цилиндрической банки.
Рассчитайте и выведите объем и полную поверхность цилиндра.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	// Ввод значений переменных с консоли
	var radius, height float64
	fmt.Print("Введите радиус: ")
	fmt.Scan(&radius) // Например, 3
	fmt.Print("Введите высоты: ")
	fmt.Scan(&height) // Например, 8

	// Объем цилиндра
	volume := math.Pi * math.Pow(radius, 2) * height

	// Площадь боковой поверхности цилиндра
	sideSurfaceArea := 2 * math.Pi * radius * height

	// Площадь основания цилиндра
	baseSurfaceArea := math.Pi * math.Pow(radius, 2)

	// Площадь полной поверхность цилиндра
	surfaceArea := sideSurfaceArea + 2*baseSurfaceArea

	// Вывод результата с округлением до сотых
	fmt.Printf("Объем цилиндрической банки: %.2f\n", volume)
	fmt.Printf("Полная поверхность цилиндрической банки: %.2f\n", surfaceArea)
}
