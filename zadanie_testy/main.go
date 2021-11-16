package main

import (
	"errors"
	"fmt"
)

func main() {

	// ilemscsieuczysz, err := beit(-5)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(ilemscsieuczysz)

	// ilemsc, err := beitbetter(-5)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(ilemsc)

	test1_spodziewam_sie_error()
	test2_spodziewam_sie_20()
	test3_spodziewam_sie_100()
}

func beit(p int) (int, error) {

	if p < 1 {
		return 0, errors.New("Nie ma takiego msc")
	}
	switch p {
	case 1:
		return 10, nil
	case 2:
		return 20, nil
	case 3:
		return 30, nil
	case 4:
		return 40, nil
	case 5:
		return 50, nil
	case 6:
		return 60, nil
	case 7:
		return 70, nil
	case 8:
		return 80, nil
	case 9:
		return 90, nil
	default:
		return 100, nil
	}
}

func beitbetter(p int) (int, error) {
	if p < 0 {
		return 0, errors.New("Nie ma takiego msc")
	}
	switch {
	case p < 10:
		return p * 10, nil
	default:
		return 100, nil
	}
}

func test1_spodziewam_sie_error() {
	want := errors.New("Nie ma takiego msc")
	_, err := beitbetter(-5)
	if err.Error() == want.Error() {
		fmt.Println("Test 1 - Mam error")
	}
}

func test2_spodziewam_sie_20() {
	want := 20
	got, err := beitbetter(2)
	if err != nil {
		fmt.Println(errors.New("funkcja zjebana"))
		return
	}
	if got == want {
		fmt.Println("Test 2 - funkcja działa prawidłowo got == want")
	} else {
		fmt.Println("funkcja zjebana got!=want")
	}
}

func test3_spodziewam_sie_100() {
	want := 100
	got, err := beitbetter(12)
	if err != nil {
		fmt.Println(errors.New("funkcja zjebana"))
		return
	}
	if got == want {
		fmt.Println("Test 3 - funkcja działa prawidłowo got == want == the best")
	} else {
		fmt.Println("funkcja zjebana got!=want")
	}
}

//create the function witch as the parameter you will pass number of month.
// number must be more than 0 (if not function will return error)

// if param will be ok, function will return percent value (from 0 - 100 int) how good you know Go language.
// you can assume that every month of learning give 10 point more of knowledge of Go.
