package tasksPantela

import "errors"

type Human struct {
	Name    string
	Age     int
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("на ноль делить нельзя! тупая ты хуйня")
	} else {
	return a/b, nil
	}
}

func ValidateAge(age int) error {
	if age >= 0 {
		return nil
	} else { 
		return errors.New("невалидный возраст")
	}
}

