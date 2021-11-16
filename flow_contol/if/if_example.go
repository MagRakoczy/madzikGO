package main

import "fmt"

func main() {
	points := map[string]int{}
	points["Magda"] = 26

	if points["Magda"] > 18 {
		fmt.Println("Pełnoletnia")
	} else {
		fmt.Println("Niepełnoletnia")
	}
	if points["Magda"] == 26 {
		fmt.Println("Magda ma 26 lat")
	}

	var buy int = 200
	if buy > 200 {
		fmt.Println("Bogactwo")
	} else if buy > 100 {
		fmt.Println("Dramat")
	} else {
		fmt.Println("Bida, oszczędzaj!!")
	}

	x := "madzik"
	czyPrawda := (x == "madzik")
	if czyPrawda {
		fmt.Println("klop")
	}

}
