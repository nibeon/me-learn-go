package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// запись в файл
	data := []byte("text to file.\nВторая строка.")
	e := ioutil.WriteFile("hello.txt", data, 0600)	
	
	if e != nil {
		fmt.Println("Невозможно создать файл!\n", e)	
	}	
	
	// чтение из файла
	file_data, err := ioutil.ReadFile("hello.txt")
	
	if err != nil {
		fmt.Println("Невозможно прочитать файл\n", err)	
	}
	
	fmt.Println(string(file_data))
	
	os.Remove("hello.txt")
}