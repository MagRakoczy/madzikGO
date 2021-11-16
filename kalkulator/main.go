package main

import (
	"errors"
	"fmt"
)

// type Wynik struct {
// 	a int
// 	b int
// }

func main() {
	// fmt.Println(add(1, 2))
	// fmt.Println(sub(1, 2))
	// fmt.Println(multi(1, 2))
	//fmt.Println(div(1, 0))
	wynikniezfunckji, err := div(1, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(wynikniezfunckji)

}

func add(a int, b int) int {
	wynikadd := a + b
	return wynikadd

}

func sub(a int, b int) int {
	wyniksub := a - b
	return wyniksub
}

func multi(a int, b int) int {
	wynik := a * b
	return wynik
}

func div(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Error: Nie można dzielić przez 0")
	}
	wynik := a / b
	return wynik, nil
}
