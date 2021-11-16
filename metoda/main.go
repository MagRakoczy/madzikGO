package main

import (
	"fmt"
)

type Cat struct {
	Name  string
	Age   int
	Sound string
}

type Dog struct {
	// Name  string
	// Age   int
	Sound string
}

type Kruczek struct {
	Name string
	Age  int
	Love string
}

func (c Cat) Speak() {
	fmt.Println(c.Sound)
}

func (c Cat) Speak3times() {
	c.Sound = fmt.Sprintf("%v! %v! %v!\n", c.Sound, c.Sound, c.Sound)
	fmt.Println(c.Sound)
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

// func Speak() {
// 	fmt.Println(d.Sound)
// }

func main() {
	Milka := Cat{Name: "Milka", Age: 2, Sound: "Miaauuu"}
	fmt.Println(Milka)
	Milka.Speak()
	Milka.Sound = "Mrau"
	Milka.Speak()
	Milka.Speak3times()
	Milka.Speak3times()

	Drops := Dog{Sound: "hauu"}
	//Drops := Dog{Name: "Drops", Age: 2, Sound: "Hauuu"}
	//Drops.Sound = "Hau"
	Drops.Speak()
	Drops.Sound = "Hauuuu"
}
