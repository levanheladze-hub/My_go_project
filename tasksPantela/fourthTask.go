package tasksPantela

import (
	"fmt"
)

func xTwo(num int) int {
	return num * 2
}

func xTwoReturnNumbers(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 * num2
}

func xError(x int) (float64, error) {
	if x <= 1 {
		return 0.0, fmt.Errorf("число меньше или равно 1")
	}
	return 1.0/float64(x), nil
}