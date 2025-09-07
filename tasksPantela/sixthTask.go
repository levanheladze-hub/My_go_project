package tasksPantela

type Address struct {
	City    string
	Country string
}

type Person struct {
	Name    string
	Age     int
	Address Address 
}

func (person *Person) changeAddress() {
    person.Address.City = "Saint Petersburg" // Изменяем оригинал
}
