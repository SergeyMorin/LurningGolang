package main

import ("fmt"
	"math")

func main(){
	var num float64
	fmt.Print("Введите число квадрат каторого необходимо найти: ")
	fmt.Scan(&num)
	fmt.Println(powerNumer(num))
}

func powerNumer(number float64)string{
	powNum := math.Pow(number, 2)
	return "Квадратом числа " + fmt.Sprint(number) + ", будет число: " + fmt.Sprint(powNum)
}