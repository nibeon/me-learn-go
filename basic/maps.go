package main

import "fmt"

func main() {
	var money map[string]int = map[string]int{
		"dollars": 1000,
		"euros":   2000,
		"apples":  3,
	}
	fmt.Println(money)

	fmt.Println(money["dollars"])

	money["dollars"] = 5000
	fmt.Println(money["dollars"])

	delete(money, "apples")
	fmt.Println(money)

	el, status := money["dollars"]
	fmt.Println(el, status)

}
