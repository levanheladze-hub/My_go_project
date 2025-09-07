package main

import (
	"fmt"
	"my_go_project/tasksPantela"
)

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

	/* //Пятое задание
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
	*/

	/* //Шестое задание
	address := Address{City: "Moscow", Country: "Russia"}
    person := &Person{Name: "Vladimir", Age: 25, Address: address} // Ссылка на address

    fmt.Println("Город до изменения:", person.Address.City)

    person.changeAddress() // Передаем указатель

    fmt.Println("Город после изменения:", person.Address.City)
	*/


	/* //Седьмое задание 	
	laptop1 := &tasksPantela.Laptop{Brand: "Lenovo", Price: 88.2} //nolint:govet
	laptop2 := &tasksPantela.Laptop{Brand: "Acer", Price: 90.4} //nolint:govet
	laptop3 := &tasksPantela.Laptop{Brand: "Asus", Price: 100.1} //nolint:govet
	laptops := []*tasksPantela.Laptop{laptop1, laptop2, laptop3} //nolint:govet

	for index, value := range laptops {
	fmt.Printf("Индекс: %d, ноутбук бренда %s цена %.1f\n", index, value.Brand, value.Price)
	}

	laptops[2] =  &tasksPantela.Laptop{Brand: "Asus", Price: 115.1} //nolint:govet

	for index, value := range laptops {
    fmt.Printf("Индекс: %d, ноутбук бренда %s цена %.1f\n", index, value.Brand, value.Price)
	}

	laptops = append(laptops, &tasksPantela.Laptop{Brand: "Nitro", Price: 70.4})
	for index, value := range laptops {
    fmt.Printf("Индекс: %d, ноутбук бренда %s цена %.1f\n", index, value.Brand, value.Price)
	}
	*/

	//Восьмое задание 
	for category, foods := range tasksPantela.Menu {
    	fmt.Printf("=== %s ===\n", category)
    	for item, price := range foods {
        	fmt.Printf("%s стоит %.2f рублей\n", item, price)
    	}
    	fmt.Println() // Пустая строка между категориями
	}
	
	value, ok := tasksPantela.Menu["Напитки"]["Бамбл"] 
    fmt.Println("Бамбл:", value, ok)

	value, ok = tasksPantela.Menu["Напитки"]["Лимонад"] 
	fmt.Println("Лимонад:", value, ok)

	value, ok = tasksPantela.Menu["На углях"]["Любой товар"]
	fmt.Println("На углях:", value, ok)

	delete(tasksPantela.Menu["Напитки"], "Лимонад")
	fmt.Println("После удаления лимонада:", tasksPantela.Menu)

	for category, foods := range tasksPantela.Menu {
    	fmt.Printf("=== %s ===\n", category)
    	for item, price := range foods {
        	fmt.Printf("%s стоит %.2f рублей\n", item, price)
    	}
    	fmt.Println() // Пустая строка между категориями
	}
}
