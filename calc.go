package main

import (
	"fmt"
)

var rightOps = map[string]bool{
	"+": true,
	"-": true,
	"/": true,
	"*": true,
}

func calc(a float64, b float64, op string) (float64, int) {
	fmt.Printf("Ваш пример: %f %s %f = ?\n", a, op, b)

	var result float64
	var code int

	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			code = 1
		}

		result = 0
	}

	return result, code
}

func main() {
	var a, b float64
	fmt.Printf("Введите число A: ")

	_, err := fmt.Scanf("%f", &a)
	if err != nil {
		fmt.Printf("Invalid data type")
	}

	fmt.Printf("Введите число B: ")

	_, err2 := fmt.Scanln("%f", &b)
	if err2 != nil {
		fmt.Printf("Invalid data type")
	}

	var op string

	fmt.Printf("Укажите операцию(/, *, +, -): ")

	fmt.Scanln(&op)

	if !rightOps[op] {
		fmt.Printf("Неизвестная операция. ")
		return
	}

	result, code := calc(a, b, op)

	if code == 1 {
		fmt.Printf("Division by zero")
		return
	}

	fmt.Printf("Результат: %f %s %f = %f", a, op, b, result)
}