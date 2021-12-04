// работа с указателями
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	x := 10
	p := &x

	fmt.Println("Значение х:", x)
	fmt.Println("Адресс x:", p)
	fmt.Println("Значение *p:", *p)

	*p = 15

	fmt.Println("Значение x после изменения *p: ", x)

}
