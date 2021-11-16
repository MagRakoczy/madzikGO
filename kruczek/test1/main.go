package main

import "fmt"

func main() {
	fmt.Println("test1")
	fmt.Println("test1")
	fmt.Print("test1\n")
	fmt.Println("tes\nt1\n\n ")
	/*komentarz
	  wieloliniowy
	  {ni chuj widać}
	*/
	//pojedyncza linia komentarza

	// var anInteger int - stworzenie zmiennej
	anInteger := 42 //krótkie stworzenie zmiennej i przypisanie wartości - deklaracja zmiennej
	fmt.Println(anInteger)

	//const anInteger int = 42 //const - zmienna const jest stała, nie można nic dodać/odjąć bo wyskoczy błąd
	//const anInteger = 42     //wartosc niejawna
	fmt.Println("\n")
	fmt.Println("HelloWorld")
	fmt.Printf("Life is beautiful")
}
