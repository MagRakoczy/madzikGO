package main

import "fmt"

func main() {
	birthday := map[string]string{
		"Magda":   "18.01.1995",
		"Kruczek": "18.07.1986",
		"Asik":    "12.12.1983",
	}
	fmt.Println(birthday)

	ages := map[string]int{}
	ages["Magda"] = 26
	ages["Kruczek"] = 35
	ages["Asik"] = 39
	fmt.Println(ages)
	fmt.Println(ages, ages["Asik"])
	delete(ages, "Asik")
	fmt.Println(ages)

}
