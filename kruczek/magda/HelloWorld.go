package main

import "fmt"

func main() {
	fmt.Println("Hello Wolrd")
	fmt.Println("Life is beutiful")
	fmt.Println("Simple String")
	fmt.Println("Hello \nWorld")

	a := 8
	b := 16
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a + b)
	SimpleString := "Hello World"
	fmt.Println(SimpleString)
	a1 := "Hello"
	b1 := "World"
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(a1 + "\n" + b1)

	fmt.Println("\\")

	kruczekdodał := "bardzo"
	SimpleString1 := ("Magda jest " + kruczekdodał + " tępa")
	SimpleString2 := ("A ja go i tak " + kruczekdodał + " kocham")
	fmt.Println(SimpleString1)
	fmt.Println(SimpleString2)

	kruczek := "kruczek"
	magda := "madzik"
	spacja := " "
	plus := "+"
	kruczek = kruczek + spacja + plus + spacja + magda

	fmt.Println(kruczek)

}
