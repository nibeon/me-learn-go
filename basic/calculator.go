package main

import (
	"fmt"
)

func main() {
	fmt.Println("Калдькулятор!")
	fmt.Println("Какое действите хотите выполнить? (+, -, *, /)")

	var action string
	fmt.Scan(&action)

	var a float64
	var b float64

	fmt.Println("Первое число: ")
	fmt.Scan(&a)

	fmt.Println("Второе число: ")
	fmt.Scan(&b)

	switch {
	case action == "+":
		fmt.Println("Сумма чисел " + fmt.Sprint(a+b))
	case action == "-":
		fmt.Println("Разница чисел " + fmt.Sprint(a-b))
	case action == "*":
		fmt.Println("Произведение чисел " + fmt.Sprint(a*b))
	case action == "/":
		fmt.Println("Деление чисел " + fmt.Sprint(a/b))
	}
}
