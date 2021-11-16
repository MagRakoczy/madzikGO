package main

import (
	"fmt"
	"time"
)

type Cat struct {
	Name   string
	Weight int
	Birth  time.Time
	Age    int
}

func main() {
	// Milka := Cat{"Milka", 2, ageofMilka}
	// fmt.Println(Milka)
	// fmt.Printf("%+v\n", Milka)

	// fmt.Printf("Name: %v\nWeight: %v kg\n", Milka.Name, Milka.Weight)

	birth := time.Date(2019, 5, 4, 00, 00, 00, 00, time.UTC)
	//fmt.Println(birth)
	currenttime := time.Now()
	diff := currenttime.Sub(birth)
	// fmt.Println(diff)
	// days := diff / 24
	days := diff.Hours() / 24
	years := days / 365
	ageofMilka := int(years)
	//fmt.Println(ageofMilka)
	Milka := Cat{Name: "Milka", Weight: 2, Birth: birth, Age: ageofMilka} //literał tworzenie zmiennej
	fmt.Println(Milka)
	//fmt.Printf("%+v\n", Milka)
	fmt.Printf("Name: %v\nWeight: %v kg\nBirthday: %v\nAge: %v lat\n", Milka.Name, Milka.Weight, Milka.Birth, Milka.Age)

	// hours := age / 60
	// fmt.Println(hours)
	// days := age / 24
	// fmt.Println()
	// years := days / 365
	// fmt.Println(years)
	// ageofMilka := int(years)
	// fmt.Println(ageofMilka)
}
