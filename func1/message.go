package main

import "fmt"

func greeting2(name, message string) string {
	return fmt.Sprintf("%s %s", message, name)
}
