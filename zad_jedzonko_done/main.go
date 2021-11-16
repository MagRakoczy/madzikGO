package main

import (
	"errors"
	"fmt"
)

type MAKRO struct {
	węgle    int
	białko   int
	tłuszcze int
	kcal     int
	zdrowe   bool
}

func main() {
	makro, err := OWOC("hamburger")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(makro.zdrowe)
	//czyzdrowe := makro.zdrowe
	if makro.zdrowe {
		fmt.Println("ZDROWE")
	} else {
		fmt.Println("NIEZDROWE NIE ŻREĆ")
	}
}

// 	czyzdrowe(makro.zdrowe)
// }
// func czyzdrowe(b bool) {
// 	if b {
// 		fmt.Println("ZDROWE")
// 	} else {
// 		fmt.Println("NIEZDROWE, NIE ŻREĆ")
// 	}
// }

func OWOC(s string) (MAKRO, error) {
	switch s {
	case "jabłko":
		return MAKRO{węgle: 1, białko: 1, tłuszcze: 1, kcal: 1, zdrowe: true}, nil
	case "pomarańcz":
		return MAKRO{węgle: 2, białko: 2, tłuszcze: 2, kcal: 2, zdrowe: true}, nil
	case "banan":
		return MAKRO{węgle: 3, białko: 3, tłuszcze: 3, kcal: 3, zdrowe: true}, nil
	case "hamburger":
		return MAKRO{węgle: 100, białko: 100, tłuszcze: 100, kcal: 1000, zdrowe: false}, nil
	default:
		return MAKRO{}, errors.New("Bląd = owoc niedozwolony")
	}
}

// ## jedzonko

// 1) stwórz funkcję aby przyjmowała owoc(string) i wyświetlała jego wartości makro (wegle, tłuszcz, białko, kacl)
//    niech funkcja obsługuje następujące owoce: jabłko, pomarancza, banan
// 2) zmodyfikuj funkcj e tak aby zwracała error jesli poda sie inne owoce niz dozowolone
// 3) stwórz strukturę makro z polami tłuszcze, węgle, białko i kalori o typie int
// 4) zmodyfikuj funkcje żeby zamiast wyswietlania wiadomosci o makro zwracała powyższa strukture
//5) dodaj do makro pole healthy bool, rozszerz nasz owoc o hamburgera
//6) po wywołaniu funckji sprAWDŹ czy dane pożywienie jest zdrowe lub nie, jeśli jest niezdrowe wyświetl komunikat: NIE ŻREĆ
