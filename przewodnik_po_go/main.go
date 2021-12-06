package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func main() {

	// rand.Seed(time.Now().UnixNano()) //ziarno utrudnia działanie

	// fmt.Println("Moją ulubioną lb jest lb: ", rand.Intn(10))

	// fmt.Printf("Teraz mam %g problemów\n",
	// 	math.Sqrt(1))
	// fmt.Println(math.Pi)

	// fmt.Println(add(4, 3))

	// names := [4]string{
	// 	"John",
	// 	"Paul",
	// 	"George",
	// 	"Ringo",
	// }
	// fmt.Println(names)

	// a := names[0:2]
	// b := names[1:3]
	// fmt.Println(a, b)

	// b[0] = "XXX"
	// fmt.Println(a, b)
	// fmt.Println(names)

	a := make([]int, 5)
	printSlice("a", a)
	// fmt.Println("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)
	// fmt.Println("b", b)

	c := b[:2]
	printSlice("c", c)
	// fmt.Println("c", c)

	d := c[2:5]
	printSlice("d", d)
	// fmt.Println("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
