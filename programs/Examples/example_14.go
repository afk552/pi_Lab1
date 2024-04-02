package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Abs(-5.67))
	fmt.Println(math.Ceil(5.67))
	fmt.Println(math.Floor(5.67))
	fmt.Println(math.Sqrt(16))
	fmt.Println(math.Pow(2, 3))
	fmt.Println(math.Sin(math.Pi / 2))
	fmt.Println(math.Log(10))
	fmt.Println(math.Max(3, 7))
	fmt.Println(math.Min(3, 7))
	fmt.Println(math.Mod(10, 3))
	fmt.Println(math.Round(5.67))
	fmt.Println(math.Trunc(5.67))
	fmt.Println(math.Inf(1))
	fmt.Println(math.Inf(-1))
	fmt.Println(math.NaN())
	fmt.Println(math.Exp(2))
	fmt.Println(math.Exp2(3))
	fmt.Println(math.Expm1(1))
	fmt.Println(math.Log10(100))
	fmt.Println(math.Log2(8))
	fmt.Println(math.Log1p(1))
	fmt.Println(math.Signbit(-5))
}
