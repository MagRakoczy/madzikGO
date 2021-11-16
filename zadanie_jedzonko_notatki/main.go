package main

import (
	"errors"
	"fmt"
)

// type Owoc struct {
// 	nazwa string
// 	carbo int
// 	whey int
// 	fat int
// 	kcal int
// }

func main() {

	fmt.Println("1) Wyświetlam owoc i jego makroskładniki:")
	// 	carbo, protein, fat, kcal := Owoc("jabłko")
	// 	fmt.Println(carbo, whey, fat, kcal)
	jabłko := "- jabłko : carbo, whey, fat, kcal"
	fmt.Println(jabłko)

	pomarańcz := "- pomarańcz: carbo, protein, fat, kcal"
	fmt.Println(pomarańcz)

	banan := "- banan : carbo, whey, fat, kcal"
	fmt.Println(banan)

	fmt.Println("###########################################")

	fmt.Println("2) Tworzę switcha: ")

	owoc := "gruszka" //mogę tu podstawić jabłko, pomarańcz, banan - i przy wartościach wyświeli mi makro, w przypadku innych wartości(nazw owoców)wyświetli mi że ten owoc jest niedowolony
	switch owoc {
	case "jabłko":
		fmt.Println("Jabłko ma węgle, białka, tłuszcze, kcal")
	case "pomarańćz":
		fmt.Println("Pomarańcz ma węgle, białka, tłuszcze, kcal")
	case "banan":
		fmt.Println("Banan ma węgle, białka, tłuszcze, kcal")
	default:
		fmt.Println("To nie jest dozwolony owoc i nie pokażę Ci makro")
	}
	fmt.Println("###########################################")

	fmt.Println("3) Za pomocą maps: ")
	owocmap := map[string]int{}
	owocmap["jabłko"] = 26
	owocmap["pomarańcz"] = 35
	owocmap["banan"] = 39
	fmt.Println(owocmap)

	fmt.Println("###########################################")

	fmt.Println("4) Za pomocą -tworzę funkcję JABŁKO: -to jest CHUJOWY SPOSÓB")
	jabłkowęgle := "węgle: 1"
	jabłkobiałka := "Białko: 1"
	jabłkotłuszcze := "Tłuszcze: 1"
	jabłkokcal := "Kcal: 1"
	JABŁKO(jabłkowęgle, jabłkobiałka, jabłkotłuszcze, jabłkokcal) //+func JABŁKO

	fmt.Println("###########################################")

	fmt.Println("5) Za pomocą -tworzę funkcję OWOC: ")
	tekst, err := OWOC("kiwi")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tekst)
}

func OWOC(s string) (string, error) {
	switch s {
	case "jabłko":
		return "jabłko ma tyle kcal", nil
	case "pomarańcz":
		return "pomarańcz ma kcal", nil
	case "banan":
		return "banan ma kcal", nil
	default:
		return "", errors.New("Bląd = owoc niedozwolony")
	}
}

func JABŁKO(jabłkowęgle, jabłkobiałka, jabłkotłuszcze, jabłkokcal string) {
	fmt.Printf("%v %v %v %v", jabłkowęgle, jabłkobiałka, jabłkotłuszcze, jabłkokcal)
}

// ## jedzonko

// 1) stwórz funkcję aby przyjmowała owoc(string) i wyświetlała jego wartości makro (wegle, tłuszcz, białko, kacl)
//    niech funkcja obsługuje następujące owoce: jabłko, pomarancza, banan
// 2) zmodyfikuj funkcj e tak aby zwracała error jesli poda sie inne owoce niz dozowolone
// 3) stwórz strukturę makro z polami tłuszcze, węgle, białko i kalori o typie int
// 4) zmodyfikuj funkcje żeby zamiast wyswietlania wiadomosci o makro zwracała powyższa strukture
// Owoc("orange")
// Owoc("apple")
// Owoc("banan")
// owoc := map[string]int{}
// owoc["jabłko"] = 1
// owoc["pomarańcz"] = 2
// owoc["banan"] = 3

// func Owoc(owoc string) {
// switch {
// case owoc["jabłko"]:
// 	fmt.Println(1, 1, 1, 1)
// case owoc["pomarańcz"] == 2:
// 	fmt.Println(2, 2, 2, 2)
// case owoc["banan"] == 3:
// 	fmt.Println(3, 3, 3, 3)
// default:
// 	fmt.Println("Ten owoc jest niedozwolony")
// }

func Owoc(owoc string) {
	if owoc == "orange" {
		fmt.Println("orange:", 1, 1, 1, 1)
		return
	}
	if owoc == "apple" {
		fmt.Println("apple:", 2, 2, 2, 2)
		return
	}
	if owoc == "banan" {
		fmt.Println("banan:", 3, 3, 3, 3)
		return
	}
}
