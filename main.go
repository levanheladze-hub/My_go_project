package main

import "fmt"

// "fmt"
// "my_go_project/tasks"


func main() {
	// fmt.Println("Hello, World!")
	// tasks.PrintZurab()
	// SumNumbers() // Первое задание
	// DayCondition() //Второе задание
	// SwitchClothes() //Второе задание
	// forUsual() //Третье задание
	// forWhile() //Третье задание
	// forRange() //Третье задание
	
	/* //Четвёртое задание
	a := xTwo(10)
	fmt.Println(a)
	b,c := xTwoReturnNumbers(3,9)
	fmt.Println(b,c)
	result, err := xError(1)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
	*/ 

	//Встроенный конструктор 
	car1 := Car{"BMW", 2000, 350.1}
	fmt.Println(car1)
	// Первичное создание уазика
	car2 := carCreation("Уазик Зураба", 2025)
	fmt.Println(car2)
	//Увеличение скорости через метод 
	car2.increasingSpeed(666)
	fmt.Println(car2)
	//Увеличение скорости через функцию
	increasingSpeed(&car2, 999)
	fmt.Println(car2)
}
