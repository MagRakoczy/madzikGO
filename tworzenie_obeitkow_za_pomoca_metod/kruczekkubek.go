package main

import "fmt"

type Kubek struct {
	napis  string
	amount int
}

func KubekDlaKruczka(napis string, amount int) Kubek {
	rezultatmetody := Kubek{}
	rezultatmetody.napis = napis
	rezultatmetody.amount = amount
	return rezultatmetody
}

func main() {
	KubekKruczek := KubekDlaKruczka("Kruczek morsuje", 1)
	fmt.Println(KubekKruczek)
	fmt.Println(dodawanie(1, 2))
}

func dodawanie(a int, b int) int {
	rezultatmetody := a + b
	return rezultatmetody
}

// func greeting(messages string, name string) string {
// 	return fmt.Sprintf("%s %s", messages, name)
