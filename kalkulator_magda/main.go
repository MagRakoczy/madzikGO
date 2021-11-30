package main

import "fmt"

func main() {

	var lb1 float64
	var lb2 float64
	var operator string

	fmt.Print("lb1:")
	fmt.Scanln(&lb1)

	fmt.Print("lb2:")
	fmt.Scanln(&lb2)

	fmt.Print("operator: +,-,*,/")
	fmt.Scanln(&operator)

	for operator == "+" {
		fmt.Printf("%f %s %f = %f", lb1, operator, lb2, lb1+lb2)
		break
	}
	for operator == "-" {
		fmt.Printf("%f %s %f = %f", lb1, operator, lb2, lb1-lb2)
	}
	for operator == "*" {
		fmt.Printf("%f %s %f = %f", lb1, operator, lb2, lb1*lb2)
	}
	for operator == "/" {
		fmt.Printf("%f %s %f = %f", lb1, operator, lb2, lb1/lb2)
	}
	// switch operator {
	// case "+":
	// 	fmt.Printf("%f %s %f = %f", lb1, operator, lb2, lb1+lb2)

	// case "-":
	// 	fmt.Printf("%f %s %f = %f", lb1, operator, lb2, lb1-lb2)

	// case "*":
	// 	fmt.Printf("%f %s %f = %f", lb1, operator, lb2, lb1*lb2)

	// case "/":
	// 	if lb2 == 0 {
	// 		fmt.Println("You cannot divide by 0")
	// 	} else {
	// 		fmt.Printf("%f %s %f = %f", lb1, operator, lb2, lb1/lb2)
	// 	}
	// default:
	// 	fmt.Println("There is not operator")
	// }
}
