package main

import "fmt"

func main() {

	var lb1 float64
	var lb2 float64
	var operator string
	for {
		fmt.Print("lb1:")
		fmt.Scanln(&lb1)

		fmt.Print("lb2:")
		fmt.Scanln(&lb2)

		fmt.Print("operator: +,-,*,/, exit")
		fmt.Scanln(&operator)

		switch operator {
		case "+":
			fmt.Printf("%f %s %f = %f\n", lb1, operator, lb2, lb1+lb2)

		case "-":
			fmt.Printf("%f %s %f = %f\n", lb1, operator, lb2, lb1-lb2)

		case "*":
			fmt.Printf("%f %s %f = %f\n", lb1, operator, lb2, lb1*lb2)

		case "/":
			if lb2 == 0 {
				fmt.Println("You cannot divide by 0")
			} else {
				fmt.Printf("%f %s %f = %f\n", lb1, operator, lb2, lb1/lb2)
			}
		case "exit":
			return
		default:
			fmt.Println("There is not operator")
		}
	}
}
