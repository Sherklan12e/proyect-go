package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operacion string
	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)
	fmt.Print("Digite el segundo número: ")
	fmt.Scan(&num2)
	fmt.Print("Digite a operacion (+, -, *, /): ")
	fmt.Scan(&operacion)
	switch operacion {
	case "+":
		fmt.Println("Resultado: ", num1+num2)

	case "-":
		fmt.Println("Resultado: ", num1-num2)
	case "*":
		fmt.Println("Resultado: ", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Println("Resultado: ", num1/num2)
		} else {
			fmt.Println("Error: Division by zero is not allowed.")
		}
	default:
		fmt.Println("Error: Operacion no valida.")
	}
}
