package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file_data, err := ioutil.ReadFile("text.txt")

	if err != nil {
		fmt.Println("Невозможно прочитать файл!\n", err)
	}

	fmt.Println(string(file_data))
}