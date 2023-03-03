package main

import "fmt"

func main(){
	var numA float64
	fmt.Print("Введите число: ")
	fmt.Scan(&numA)

	switch(numA) {
        case 1: 
            fmt.Println("Понедельник")
        case 2: 
            fmt.Println("Вторник")
        case 3: 
            fmt.Println("Среда")
		case 4: 
            fmt.Println("Четверг")
        case 5: 
            fmt.Println("Пятница")
        case 6: 
            fmt.Println("Суббота")
		case 7:
			fmt.Println("Воскресение")
		default: 
            fmt.Println("Такого дня недели нет")
    }
}