package main

import "fmt"

func test2(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func test(somefunction func(int) int) {
	fmt.Println(somefunction(25))
}

func main() {
	// использование анонимных функций
	a := func(x int, y int) int {
		return x + y
	}
	sum := a(3, 4)
	fmt.Println(sum)

	// передача параметров из внешней функции
	square := func(x int) int {
		return x * x
	}
	test(square)

	// еще один варпиант использование функции
	test2("Hello")()
}
