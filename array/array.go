package main

import "fmt"

func main() {
	//names := [3]string{"Madzik", "Kruczek", "Love"}
	var names [4]string
	names[1] = "Madzik"
	names[2] = "Kruczek"
	names[0] = "Love"
	//names[3] = "is nil"
	fmt.Println(names)
	fmt.Println(names[3])
	fmt.Println(names[3] == "") //- pusty znak

	//profil := [4]string{"Magdalena,", "Rakoczy,", "18.01.2021,", "Świętochłowice"}
	//fmt.Println(profil)

	// var profil [4]string
	// profil[1] = "Magdalena"
	// profil[2] = "Rakoczy"
	// profil[3] = "18.01.1995"
	// profil[0] = "Świętochłowice"
	// fmt.Println(profil)

	imię := "Magdalena"
	nazwisko := "Rakoczy"
	dataur := "18.01.1995"
	miejsceur := "Świętochłowice"
	var profil [4]string
	profil[0] = imię
	profil[1] = nazwisko
	profil[2] = dataur
	profil[3] = miejsceur
	fmt.Println(profil)

}
