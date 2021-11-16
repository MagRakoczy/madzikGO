package main

import "fmt"

func main() {
	var p *int
	var x int = 42
	p = &x
	if p != nil {
		fmt.Println("Wartość p:", *p)
		fmt.Println("Adres p:", p)
		fmt.Println("Adres p", &x)
	} else {
		fmt.Println("*p is nil")
	}

	var x1 float32 = 42.64
	p1 := &x1
	fmt.Println("wartość x1 :", *p1)
	*p1 = *p1 / 31
	fmt.Println("x1: :", x1)

}
