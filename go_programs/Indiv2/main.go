/*
Индивидуальное задание 2 (8). Даны основания равнобедренной
трапеции и угол при большем основании. Найти площадь трапеции.
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
