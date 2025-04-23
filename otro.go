package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum of first 10 numbers:", sum)
	a := 0
	for a < 10 {
		a += a + 1
		fmt.Println("This will run forever", a)
	}
	b := 1
	if b > 4 {
		fmt.Println("b is 4", b)
		b++
	}
	var ala [3]string
	ala[0] = "ala"
	ala[1] = "bala"
	ala[2] = "bala"
	fmt.Println(ala)
	primos := [3]string{"uno", "dos", "tres"}
	fmt.Println(primos)

}
