package main

import "fmt"

func main() {
	fmt.Println("Simple String")
	fmt.Println(`
This is multi line
String, that can also contain "quotes"
`)
	fmt.Println("Also this is \n new line \n and new line")
	fmt.Println(`Natomiast pojedynczy cudzysłów i \n nie działają`)
	fmt.Println(`Musiałaby to zrobić tak 
	\n
	i po prostu enterem podzielić`)
}
