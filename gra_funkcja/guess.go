package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//var n int //random nb
	//0 < n < 100

	czyZgadlem := fgra()
	if czyZgadlem {
		fmt.Println("super")
	} else {
		fmt.Println("nie zgadłeś, brak więcej prób")
	}
	// if err != nil {
	// 	fmt.Println("błąd")
	// } else if t == 0 {
	// 	fmt.Println("Brak prób", 0)
	// } else {
	// 	fmt.Println(wynik, t)
	// }
}

func fgra() bool {
	rand.Seed(time.Now().UnixMilli())
	n := rand.Intn(100)
	var liczbalosowa int
	//var t int //ilość prób
	for t := 10; t > 0; {
		t--
		fmt.Print("Podaj losową liczbę: ")
		fmt.Scanln(&liczbalosowa)
		if liczbalosowa < 0 {
			fmt.Println("Nie moze byc minus", "zostało prób: ", t)
		}
		if liczbalosowa < n {
			fmt.Println("Ups, podana liczba jest za mała, zostało prób: ", t)
		}
		if liczbalosowa > n {
			fmt.Println("Ups, podana lb jest za duża, zostało prób: ", t)
		}
		if liczbalosowa == n {
			// fmt.Println("BRAWO, zgadłeś")
			return true
		}
	}
	return false
}
