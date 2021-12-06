package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//var n int //random nb
	//0 < n < 100
	rand.Seed(time.Now().UnixMilli())
	n := rand.Intn(100)
	//fmt.Println(n)
	var liczbalosowa int
	//var t int //ilość prób
	for t := 10; t > 0; {
		fmt.Print("Podaj losową liczbę: ")
		fmt.Scanln(&liczbalosowa)
		t--
		if t == 0 {
			fmt.Println("Brak prób.\nSzukana lb to: ", n)
			return
		}
		switch {
		case liczbalosowa < n:
			fmt.Println("Ups, podana liczba jest za mała, zostało prób: ", t)
		case liczbalosowa > n:
			fmt.Println("Ups, podana lb jest za duża, zostało prób: ", t)
		case liczbalosowa == n:
			fmt.Println("BRAWO, zgadłeś")
			return
		default:
			fmt.Println("hę")
		}
	}
}
