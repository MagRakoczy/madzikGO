package main

import "fmt"

func main() {
	// floor := map[string]int{}
	// floor["winda1"] = 5 //piętro

	// switch {
	// case 1 == 1:
	// 	fmt.Println("Go on first floor")
	// case floor["winda1"] == 2:
	// 	fmt.Println("Go on second floor")
	// case floor["winda1"] == 3:
	// 	fmt.Println("Go on third floor")
	// case floor["winda1"] == 4:
	// 	fmt.Println("Go on fourth floor")
	// case floor["winda1"] == 5:
	// 	fmt.Println("Go on fifth floor")
	// default: //domyślny
	// 	fmt.Println("Naciśnij przycisk")
	// }

	// switch floor["winda1"] {
	// case 1, 2, 3, 4:
	// 	fmt.Println("Piętro otwarte")
	// case 5:
	// 	fmt.Println("Piętro zastrzeżone")

	// // }
	kubek := map[string]int{}
	kubek["kawa"] = 6

	switch {
	case kubek["kawa"] == 1:
		fmt.Println("Cappucinovelove")
	case kubek["kawa"] == 2:
		fmt.Println("Lattelove")
	case kubek["kawa"] == 3:
		fmt.Println("Black Coffee")
	case kubek["kawa"] == 4:
		fmt.Println("Flat white")
	case kubek["kawa"] == 5:
		fmt.Println("Do roboty!")
	default:
		fmt.Println("A weź spierdalaj")
	}
	kubek1 := map[string]int{}
	kubek1["cappucino"] = 1
	switch kubek["cappucino"] {
	case 2, 3, 4, 5, 6:
		fmt.Println("To nie cappucino")
	case 1:
		fmt.Println("Leeeej")
	}
	cappucino := 1
	if cappucino == 1 {
		fmt.Println("Leeeej")
	} else {
		fmt.Println("To nie cappucino")
	}
	cappucinoo := 1

	switch cappucinoo {
	case 1:
		fmt.Println("leeeej")
	default:
		fmt.Println("To nie cappucino")
	}

}
