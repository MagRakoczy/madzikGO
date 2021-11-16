package main

import "fmt"

func main() {

	var myInt int = 16
	var val, ok = "hello", true

	fmt.Println("myInt is: ", myInt)

	fmt.Println(val)
	fmt.Println(ok)

	// //var x, y, z = "0", "1", "2"
	// fmt.Println("zmienna x ma wartość: ", x)
	// fmt.Println("zmienna y ma wartość: ", y)
	// fmt.Println("zmienna z ma wartość: ", z)
	// fmt.Println(x + y + z)
	a := 2
	fmt.Println(a * 2)
	var b int = 2
	fmt.Println(b)
	fmt.Println(a * b)

	// var x, y, z = 0, 1, 2
	// fmt.Println("zmienna x ma wartość: ", x)
	// fmt.Println("zmienna y ma wartość: ", y)
	// fmt.Println("zmienna z ma wartość: ", z)
	// fmt.Println(x + y + z)

	// var x, y, z float64 = 0, 1, 2.5 //jeśli zmienna ma wartość po przecinku - trzeba dodać float64
	// fmt.Println("zmienna x ma wartość: ", x)
	// fmt.Println("zmienna y ma wartość: ", y)
	// fmt.Println("zmienna z ma wartość: ", z)
	// fmt.Println(x + y + z)

	var x, y, z = "madzik", "kocha", "kruczka"
	fmt.Println("zmienna x ma wartość: ", x)
	fmt.Println("zmienna y ma wartość: ", y)
	fmt.Println("zmienna z ma wartość: ", z)
	fmt.Println(x + " " + y + " " + z)

	var name string = "madzika"
	fmt.Println(name)
}
