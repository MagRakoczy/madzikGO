package kilkaplikowwfolderze

import "fmt"

func Imienazwisko(imie string, nazwisko string) string {
	result := fmt.Sprintf("%s %s", imie, nazwisko)
	return result
}
