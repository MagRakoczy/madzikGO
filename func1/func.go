package main

import "fmt"

func main() {
	message := "Hello Magda - pierwsza linijka w funkcji main"
	fmt.Println(message)
	fmt.Println("*******************")
	messages := greeting("Hello", "Magda")
	fmt.Println(messages) //+func greeting
	fmt.Println("*******************")

	name1 := "Magda"
	name2 := "Kruczek"
	name3 := "Milcia"
	names := greetings(name1, name2, name3)
	fmt.Println(names) //+func greetings

}

func greeting(messages string, name string) string {
	return fmt.Sprintf("%s %s", messages, name)
}

func greetings(name1 string, name2 string, name3 string) string {
	//to samo mogę zapisac: (name1,name2,name3 string) string
	return fmt.Sprintf("%s %s %s", name1, name2, name3)
	//gdy Sprint bez f - bez info że spodziewa się stringa
	//return fmt.Sprint(name1, name2, name3) - ale wyjdzie bez spacji
}
