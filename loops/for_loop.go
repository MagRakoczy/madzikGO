package main

import "fmt"

func main() {
	// for i := 0; i <= 100; i = i + 10 { // = i++
	// 	fmt.Println(i)
	// }
	// fmt.Println("****")
	// for i1 := 1; i1 <= 100; i1 = i1 + 10 { // = i++
	// 	fmt.Println(i1)
	// }
	// fmt.Println("****")
	a := 0
	for a <= 20 {
		if a%2 == 0 {
			a++
			continue
		} else if a == 11 {
			break
		}
		fmt.Println("count", a)
		a++
	}

	// ages := map[string]int{}
	// ages["magda"] = 26
	// ages["kruczek"] = 35
	// ages["asik"] = 39
	// for name, age := range ages {
	// 	switch age {
	// 	case 1, 2, 3, 7, 11, 19:
	//fmt.Println(fmt.Sprintf("%s jego wiek jest lb pierwszą", name))
	// 	case 26:
	// 		fmt.Println(name, "ma 26 lat")
	// 	case 39:
	// 		fmt.Println(name, "jest najstarsza")
	// 	default:
	// 		fmt.Println(fmt.Sprintf("%s i tak cię kocham", name))
	// 	}

}
