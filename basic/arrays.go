package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		names := [3]string{"John", "Jordan", "Kate"}
		fmt.Println(names, len(names))

		for i := 0; i < len(names); i++ {
			fmt.Println(names[i])
		}
	*/

	marks := [5]float64{11, 10, 12, 8, 11} // массив из 5 элементов
	var sum float64 = 0

	for i := 0; i < len(marks); i++ {
		sum += marks[i]
	}

	var result float64 = sum / float64(len(marks))
	fmt.Println(math.Round(result))
}
