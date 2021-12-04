package main

import "fmt"

func main() {
	var name string
	var age uint8
	fmt.Println("What is your namne?")
	fmt.Scan(&name)	
	fmt.Println("Hello " + name + "!")
	fmt.Println("How old are you?")
	fmt.Scan(&age)
	fmt.Println("You arte " + fmt.Sprint(age) + " years!")
	
}
