package main

import "fmt"

func main() {
	messages := greeting("Hello", "Magda")
	fmt.Println(messages)
}

// func greeting(messages string, name string) string {
// 	return fmt.Sprintf("%s %s", messages, name)
// }
// func greeting(name, message string) (salut string) {
// 	salut = fmt.Sprintf("%s %s", message, name)
// 	return
// } = poprawny zapis w go, ale nie uzywać
func greeting(messages string, name string) string {
	result := fmt.Sprintf("%s %s", messages, name)
	return result //najczęściej uzywane
}
