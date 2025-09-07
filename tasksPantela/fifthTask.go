package tasksPantela 

// import (
// 	"fmt"
// )

type Car struct {
	Brand string
	Year int 
	Speed float64
}

func carCreation(brand string, year int) Car {
	car := Car{brand, year, float64(200)}
	return car
}

func carCreationWithSpeed(brand string, year int, speed float64) Car {
	car := Car{brand, year, speed}
	return car
}

func (car *Car) increasingSpeed(speed float64) {
	car.Speed = speed
}

func increasingSpeed(car *Car, speed float64) {
	car.Speed = speed
}