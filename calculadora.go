package main

import (
	"fmt"
)

func main() {
	a := 0
	for a < 1 {

		var num1, num2 float64
		var operacion string
		fmt.Print("Digite o primeiro número: ")
		fmt.Scan(&num1)
		fmt.Print("Digite el segundo número: ")
		fmt.Scan(&num2)
		v, err := fmt.Scanf("%f", &num1, &num2)
		if num1 == 0 && num2 == 0 || v != 1 && err == nil || v != 1 && err == nil {
			fmt.Println("Error: No se puede realizar operaciones con cero.")
			break
		}
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
			a = 1
		}
	}

}
