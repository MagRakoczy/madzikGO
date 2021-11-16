package main

import (
	"errors"
	"fmt"
)

func main() {
	samochodkruczka := car{bag: 10}
	ileprzejedzie, err := samochod(samochodkruczka)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ileprzejedzie)
	test1()
}

type car struct {
	bag int //litry w baku
}

func samochod(c car) (int, error) {
	if c.bag < 0 {
		return 0, errors.New("Nigdzie nie pojedzie, brak benzyny w baku")
	}
	ileprzejedzie := c.bag * 10
	return ileprzejedzie, nil
}

func test1() {
	want := 100
	samochodkruczka := car{bag: 10}
	got, err := samochod(samochodkruczka)
	if err != nil {
		fmt.Println(errors.New("funkcja zjebana"))
		return
	}
	if got == want {
		fmt.Println("funkcja działa prawidłowo")
	}
}

//10km = 1 L
