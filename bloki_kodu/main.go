package main

import "fmt"

func main() {
	a := 26.126
	b := "lat"
	c := 4.54151515
	fmt.Printf("Magda ma: %0.2f %s\n", a, b)
	d := a + c
	fmt.Printf("%.2f\n", d)
	fmt.Println("************************")

	fA(3, 2)
	fB(3, 3)
	fC("Magda", "kruczek")
	fmt.Println(double(4.2))
	fmt.Println(int(double(4.2)))
	fmt.Println(eheh("Magda"))
}
func fA(a int, b int) {
	fmt.Println(a + b)
}

func fB(a int, b int) {
	fmt.Println(a * b)
}

func fC(a string, b string) {
	fmt.Println(a + " + " + b)
}

func double(number float64) float64 {
	return number * 2
}
func eheh(imie string) string {
	return imie * 2
}
