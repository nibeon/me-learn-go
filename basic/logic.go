package main

import "fmt"

func main() {
	a := 3
	b := 10

	if a < 1 || b < 5 {
		fmt.Println("TRUE!")
	} else {
		fmt.Println("False")
	}
}
