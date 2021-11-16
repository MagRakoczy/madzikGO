package main

import "fmt"

func main() {
	fmt.Println("1 Greater than 2 (1>2): ", 1 > 2)
	fmt.Println("2 Less than 3 (2<3): ", 2 < 3)
	fmt.Println("czy 2>=2", 2 >= 2)
	fmt.Println("2<=2", 2 <= 2)
	fmt.Println("Equaivalent - równowartość 2.0==2:", 2.0 == 2)
	fmt.Println("Not Equivalent - Nierówna się 2.0 != 2.0:", 2.0 != 2.0)
	fmt.Println("Not Equivalent 2,0!=2,1: ", 2.1 != 2.0)
	//nil - odpowiada za pustą wartość

	var x int = 1
	fmt.Println(x)

	var y int //w domyśle = 0
	fmt.Println(y)
	fmt.Println("x<y", x < y)
}
