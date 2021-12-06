package main

import (
	"fmt"
)

func main() {

	// 	var price int = 3
	// 	var tax float64 = 2.5
	// 	var total float64 = float64(price) + tax
	// 	fmt.Println(total)
	// 	var available int = 1
	// 	if float64(available) > total {
	// 		fmt.Println("tak, zmieszcze się")
	// 	} else {
	// 		fmt.Println("za mało hajsu")
	// 	}
	// 	fmt.Println("go run")

	// fmt.Println("Podaj wynik:")
	// reader := bufio.NewReader(os.Stdin)
	// input, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// input = strings.TrimSpace(input)
	// grade, err := strconv.ParseFloat(input, 64)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var status string
	// if grade >= 60 {
	// 	status = "zaliczenie"
	// } else {
	// 	status = "brak zaliczenia"
	// }
	// fmt.Println("wynik", grade, "oznacza", status, ".")

	var lbpunktow float64

	fmt.Print("Liczba punktów z testu: ")
	fmt.Scanln(&lbpunktow)

	fmt.Println(lbpunktow)

	if lbpunktow == 100 {
		fmt.Println("Doskonale")
	}
	if lbpunktow > 60 {
		fmt.Println("Zdane")
	} else {
		fmt.Println("Nie zdane")
	}

	//??
	// i, err := strconv.Atoi(input)
	// if err != nil {
	// 	fmt.Println(err, "dupa")
	// }
	// if i >= 100 {
	// 	fmt.Println("doskonale")
	// } else {
	// 	fmt.Println("dupa")
	// }
	// if i < 100 {
	// 	fmt.Println("jest git")
	// } else {
	// 	fmt.Println("dupa")
	// }
}
