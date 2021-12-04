package main

import "fmt"

func sum(a int, b int) {
	fmt.Println(a + b)
}

func sum2(a int, b int) int {
	result := a + b
	return result
}

func math_func(a int, b int) (int, int, int, int) {
	sum := a + b
	sub := a - b
	mult := a * b
	div := a / b
	return sum, sub, mult, div
}

// другой вариант этой же функции
func variant_math_func(a int, b int) (sum int, sub int, mult int, div int) {
	sum = a + b
	sub = a - b
	mult = a * b
	div = a / b
	return
}

func main() {
	sum(5, 9)

	value := sum2(3, 4)
	fmt.Println(value)

	s, su, m, d := math_func(3, 5)
	fmt.Println(s, su, m, d)

	s2, su2, m2, d2 := variant_math_func(3, 5)
	fmt.Println(s2, su2, m2, d2)
}
