package main

import "fmt"

func main() {
	// firstname := "Magdalena"
	// lastname := "Rakoczy"
	// full := firstname + lastname
	// fmt.Println(full)
	// fmt.Println("*****************")
	n1, l1 := fullName("Magdalena", "Rakoczy") //ta opcja jest lepsza
	//n1, l1 := fullNameNakedReturn("Magdalena", "Rakoczy")
	fmt.Println(n1, l1, "-FullName, ma tyle znaków")
}

func fullName(firstname string, lastname string) (string, int) { //ta opcja jest lepsza
	full := firstname + " " + lastname
	lenght := len(full)
	return full, lenght
}

// func fullNameNakedReturn(firstname, lastname string) (full string, lenght int) {
// 	full = firstname + " " + lastname
// 	lenght = len(full)
// 	return
// }
