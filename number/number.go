package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Dodawanie - Addition:", 1+3)
	fmt.Println("Odejmowanie -Substraction", 4-1)
	fmt.Println("Dzielenie - Division:", 14/4)
	fmt.Println("Mnożenie - Multiplication:", 3*4)

	fmt.Println(1+3, " - dodawanie")

	x := (10.0 / 4)
	fmt.Println(x)
	zmiananacałkowita := int(10 / 4)
	fmt.Println(zmiananacałkowita)
	zmianaXnacałkowita := int(x)
	fmt.Println(zmianaXnacałkowita)

	fmt.Println("potęgowanie:", math.Pow(2, 2))
	var y float64 = 2
	fmt.Println(y)
	//nowa wartość dla x:
	x = math.Pow(y, 2)
	fmt.Println("potęga lb y:", y)

	var z float64 = 4
	y = math.Pow(z, 2)
	fmt.Println("potęga lb z:", z)

	fmt.Println("potęgowanie 10^2:", math.Pow(10.0, 2))

}
