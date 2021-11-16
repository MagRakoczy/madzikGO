package main

import (
	"errors"
	"fmt"
)

func main() {

	wypłata := WYPŁATA(45000)
	fmt.Println(wypłata)

	//fmt.Println(Targety(-50))
	target, err := Targety(50)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(target)

	fmt.Println(Targety(10))
	wciele_test2_spodziewam_sie_error()
	test3000()
	test4000()
	test5000()
	test6000()
}

func wciele_test1_spodziewam_sie_2000() {

	want := 2000
	// got = wynnik z funkcji Target przy wart 10
	got, err := Targety(10)
	if err != nil {
		fmt.Println(errors.New("funkcja zjebna"))
		return
	}
	if got == want {
		fmt.Println("funkcja działa prawidłowo")
	}
}

func wciele_test2_spodziewam_sie_error() {

	want := errors.New("Error, value is below zero")
	// got = wynnik z funkcji Target przy wart 10
	_, err := Targety(-10)
	if err.Error() == want.Error() {
		fmt.Println("funkcja dziala prawidlowo")
	}
}

func test3000() {
	want := 3000
	// got = wynnik z funkcji Target przy wart 10
	got, err := Targety(30)
	if err != nil {
		fmt.Println(errors.New("funkcja zjebana"))
		return
	}
	if got == want {
		fmt.Println("funkcja działa prawidłowo")
	}
}
func test4000() {
	want := 4000
	// got = wynnik z funkcji Target przy wart 10
	got, err := Targety(50)
	if err != nil {
		fmt.Println(errors.New("funkcja zjebana"))
		return
	}
	if got == want {
		fmt.Println("funkcja działa prawidłowo")
	}
}
func test5000() {
	want := 5000
	// got = wynnik z funkcji Target przy wart 10
	got, err := Targety(80)
	if err != nil {
		fmt.Println(errors.New("funkcja zjebana"))
		return
	}
	if got == want {
		fmt.Println("funkcja działa prawidłowo")
	}
}
func test6000() {
	want := 6000
	// got = wynnik z funkcji Target przy wart 10
	got, err := Targety(101)
	if err != nil {
		fmt.Println(errors.New("funkcja zjebana"))
		return
	}
	if got == want {
		fmt.Println("funkcja działa prawidłowo")
	}
}

func WYPŁATA(p int) int {

	golaWyplata := 2000

	switch {
	// case p < 30000:
	// 	return golaWyplata + 1000
	case (30000 < p) && (p <= 60000):
		return golaWyplata + 1000
	case (60000 < p) && (p <= 90000):
		return golaWyplata + 2000
	case p > 90000:
		return 5000
	default:
		return golaWyplata
	}
}

func Targety(p int) (int, error) {

	podstwypłata := 2000
	switch {
	// case p <= 20:
	// 	return podstwypłata
	case (p > 20) && (p <= 40):
		return podstwypłata + 1000, nil
	case (p > 40) && (p <= 60):
		return podstwypłata + 2000, nil
	case (p > 60) && (p <= 100):
		return podstwypłata + 3000, nil
	case p > 101:
		return podstwypłata + 4000, nil
	case p < 0:
		return 0, errors.New("Error, value is below zero")
	default:
		return podstwypłata, nil
	}
}

// stwórz funkcje która przyjmnie paramtert wykonanego planu sprzedazy i zwroci wartosc wypłaty dla pracowika:

// progi:
//  do 30 000 -> wypłata 2000
//  od 30 001 do 60 000 -> wypłata 3000
//  od 60 001 do 90 000 -> 4000
//  od 90 000 -> wypłata 5000

// now change the value of parameter, now func take amount of money,
// but i want that funcion can take the percent (int) of whole goal... ex. your monthly goal is 100 000,
// and for now you have 50 000 so to the function you will put 50
// yours cases:
// <20% -> 2000
// 21%-40 ->3000
// 41-60 -> 4000
// 61-100 -> 5000
// 101 -> 6000
