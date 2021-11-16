package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"input/kilkaplikowwfolderze"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Twoje imię : ")
	name, _ := reader.ReadString('\n') //- zniwelowanie spacji
	name = strings.TrimSpace(name)

	fmt.Print("Twoje nazwisko: ")
	nazwisko, _ := reader.ReadString('\n')
	nazwisko = strings.TrimSpace(nazwisko)

	imienazwisko := kilkaplikowwfolderze.Imienazwisko(nazwisko, name)
	fmt.Println(imienazwisko)
}
