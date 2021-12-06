package main

import "fmt"

// func reverseArray(a []int32) []int32 {
// 	for i := 0; i++ {
// 		j := len(a) - i - 1
// 		a[i], a[j] = a[j], a[i]
// 	}
// 	return a

func main() {
	// fmt.Printf("%v\n", reverseArray([]int{1, 2, 3}))
	// fmt.Printf("%v\n", reverse([]int{1, 2, 3, 4}))

	a := []int{1, 2, 3, 23}
	fmt.Println(reverseArray(a))
}

func reverseArray(a []int) []int {

	result := make([]int, len(a))

	for i, v := range a {
		result[len(a)-1-i] = v
	}

	return result
}
