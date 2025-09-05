package main

import (
	"fmt"
	"math/rand"
)

func DayCondition() {
	isSunny := rand.Intn(2) == 0
	isWeekend := rand.Intn(2) == 0
	if isSunny && isWeekend {
		fmt.Println("Идеальный день для прогулки")
	} else if isSunny && !isWeekend {
		fmt.Println("Погода хорошая, но нужно поработать")
	} else if !isSunny && isWeekend {
		fmt.Println("Можно остаться дома и отдохнуть")
	} else {
		fmt.Println("Рабочий день с плохой погодой")
	}
}

func SwitchClothes() {
	temp := rand.Intn(51)
	switch {
	case temp <= 15:
		fmt.Println("Температура: ", temp, ". Надень куртку, яички отморозишь")
	case temp <= 30:
		fmt.Println("Футболку надень, яичкам норм будет")
	case temp <= 45:
		fmt.Println("Можешь не переживать, яичница и так будет")
	default:
		fmt.Println("Дома сиди, иначе охуеешь")
	}
}
