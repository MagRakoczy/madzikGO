package main

import (
	"errors"
	"fmt"
)

// func main() {

// 	czymozejechac, komentart := KruczekJedzieDo("Tarnowskie Góry")
// 	fmt.Println(czymozejechac, komentart)
// }

// func KruczekJedzieDo(miasto string) (bool, string) {
// 	if miasto == "Tarnowskie Góry" {
// 		return true, "A weź spierdalaj"
// 	} else {
// 		return false, "vdmvsv"
// 	}
// }

func main() {
	czypisałeśdomagdy, tekst := Wiadomość("nie")
	fmt.Println(czypisałeśdomagdy, tekst)

	czyodpisałeśmadzi, err := SMS("nie")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(czyodpisałeśmadzi)

}

func Wiadomość(sms string) (bool, string) {
	if sms == "tak" {
		return true, "A weź spierdalaj"
	}
	return false, "Kc"
}

func SMS(sms string) (string, error) {
	if sms == "tak" {
		return "A weź spierdalaj", errors.New("grr")
	}
	return "Jak nie to super, Kc", nil
}
