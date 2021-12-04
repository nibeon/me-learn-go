package main

import (
	"fmt"
	"sort" // функции сортировки
)

func main() {
	slice := []int{3, 1, 2, 5, 7, 4}
	slice = append(slice, 0) // добавление элементов
	slice[0] = 11
	sort.Ints(slice) // сортировка интов

	fmt.Println(slice)

	slice2 := []string{"v", "b", "a", "c", "d", "z"}
	fmt.Println(slice2)
	sort.Strings(slice2) // сортировка строк
	fmt.Println(slice2)
}
