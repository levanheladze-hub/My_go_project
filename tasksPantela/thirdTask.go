package tasksPantela

import (
	"fmt"
	"time"
)

func forUsual() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func forWhile() {
	var counter int = 0
	for {
		counter++
		fmt.Println(counter)
		time.Sleep(1 * time.Second)
		if counter == 5 {
			fmt.Println("Досчитали до 5")
			break
		}
	}
}

func forRange() {
	numbers := []int{10, 20, 30, 40, 50}

	for index, value := range numbers {
		fmt.Printf("Индекс: %d, Значение: %d\n", index, value)
	}

	for _, value := range numbers {
		fmt.Printf("Значение: %d\n", value)
	}

	for index := range numbers {
		fmt.Printf("Индекс: %d\n", index)
	}
}
