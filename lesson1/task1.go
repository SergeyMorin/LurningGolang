package main

import ("fmt"
	"math")

func main(){
	var numA float64
	fmt.Print("Введите число A: ")
	fmt.Scan(&numA)
	var numB float64
	fmt.Print("Введите число B: ")
	fmt.Scan(&numB)

	if math.Pow(numB, 2) == numA {
		fmt.Print("Число A является квадратом числа B")
	} else {
		fmt.Print("Число A нет является квадратом числа B")
	}	
}