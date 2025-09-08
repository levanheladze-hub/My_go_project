package tasksPantela

import "fmt"

type Movable interface {
	move(s string) string
	Move() string
}

type Car struct{
	Brand string
	Speed float64
}

type Person struct{
	Name string
	Age int
	Speed float64
}

func (car Car) move(s string) string {
	return fmt.Sprintf("%s машина %s едет со скоростью %f км/ч", s, car.Brand, car.Speed)
}

func (car Car) Move() string {
	return fmt.Sprintf("Машина %s едет со скоростью %f км/ч", car.Brand, car.Speed)
}

func (person Person) move(s string) string {
	return fmt.Sprintf("%s человек возрастом %d лет идёт со скоростью %f км/ч", s, person.Age, person.Speed)
}

func (person Person) Move() string {
	return fmt.Sprintf("Человек возрастом %d лет идёт со скоростью %f км/ч", person.Age, person.Speed)
}

func PrintMovement(m Movable) {
	fmt.Println(m.Move())
}