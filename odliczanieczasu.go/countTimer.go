package main

import (
	"fmt"
	"time"
)

func main() {

	currenttime := time.Now()
	//podana data to data urodzin rakoczy
	birthtime := time.Date(1995, 1, 18, 00, 00, 00, 00, time.UTC)
	diff := currenttime.Sub(birthtime)
	days := diff.Hours() / 24
	years := days / 365
	ageOfPerson := int(years)

	// if ageOfPerson >= 18 {
	// 	fmt.Println("Pełnoletni")
	// } else {
	// 	fmt.Println("Niepełnoletni")
	// }

	if ageOfPerson < 18 {
		fmt.Println("Niepełnoletni")
		return
	}

	fmt.Println("Jestem pełnoletnia i jestem zajebista")

	fmt.Println("###################################\n")
	fmt.Println("\n###################################\n")

	var date time.Time = currenttime
	fmt.Println(date)

	// var day int = 18
	// var month int = 01
	// var year int = 1995
	// birth := (day + month + year)
	//fmt.Println("- to nie jest dobrze")
	day := 18
	month := 01
	year := 1995
	birth := fmt.Sprint(day, ".", month, ".", year)
	fmt.Println(birth)
	fmt.Println("- to nie jest dobrze")
	// var birht1 int = (day, month, year) //-czemu to jest podkeślone i nie działa?
	// fmt.Println(birth1)
	fmt.Println("###################################\n")
	fmt.Println("slice:")
	var birth1 [3]int
	birth1[0] = day
	birth1[1] = month
	birth1[2] = year
	fmt.Println(birth1)
	fmt.Println("###################################")
	fmt.Println("array:")
	var birth3 = []string{"18", ".", "01", "1995"}
	fmt.Println(birth3)
	birth4 := [3]int{day, month, year}
	fmt.Println(birth4)
	fmt.Println("###################################")
	fmt.Println("map:")
	birth2 := map[int]string{}
	birth2[day] = "/"
	birth2[month] = "/"
	birth2[year] = "/"
	fmt.Println(birth2)
	fmt.Println("###################################\n")

}
