package main

import "fmt"
import "math"

func main() {
	
	var a float64
	var b float64
	var c float64
	
	fmt.Println("Квадратное уравнение")
	
	fmt.Println("Введите a: ")
	fmt.Scan(&a)
	
	fmt.Println("Введите b: ")
	fmt.Scan(&b)
	
	fmt.Println("Введите c: ")
	fmt.Scan(&c)
	
	D := (b*b) - 4*(a*c) // дискриминант
	
	fmt.Println("Дискриминант = " + fmt.Sprint(D))
	
	if D > 0 {
		var x1 float64
		var x2 float64
		
		x1 = (-b + math.Sqrt(D)) / (2*a)
		x2 = (-b - math.Sqrt(D)) / (2*a)
	
		fmt.Println("Ваше уравнение имеет 2 корня\n D = " + fmt.Sprint(D))
		fmt.Println("X1: " + fmt.Sprint(x1) + "\n X2: " + fmt.Sprint(x2))
	}  else if D == 0 {
		var x float64
		
		x = (-b) / (2*a)
			
		fmt.Println("Ваше уравнение имеет 1 корень \n D = 0")
		fmt.Println("X: " + fmt.Sprint(x))
	} else {
		fmt.Println("Ваше уравнение не имеет корней. \n D < 0")
	}
}
