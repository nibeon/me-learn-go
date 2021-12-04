// продвинутый перебор массивов и срезов

package main

import "fmt"

func main() {
	slice := []int{1, 4, 5, 2, 6, 9, 10}

	for _, el := range slice {
		fmt.Printf("%d\n", el)
	}
}
