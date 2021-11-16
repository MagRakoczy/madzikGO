package main

import "fmt"

func main() {
	//names := []string{"Madzik"}
	names := []string{"1 element"}
	names = append(names, "+ 2 element")
	fmt.Println(names)
	names = append(names, "+ 3 element")
	fmt.Println(names)

	names1 := make([]string, 4)
	names1[1] = "Madzik"
	names1[2] = "Love"
	names1[3] = "Kruczek"
	fmt.Println(names1)

	names2 := []string{"to"}
	names2 = append(names2, "jest")
	names2 = append(names2, "skomplikowane")
	fmt.Println(names2)
	names2 = append(names2, "LOL")
	fmt.Println(names2)
	fmt.Println("tylko 4 element =", names2[3])
	fmt.Println("czy 4 elem. names to LOL? - ", names2[3] == "LOL")
	fmt.Println("czy 4 elem. names to 0? - ", names2[3] == "0")

}
