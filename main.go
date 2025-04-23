package main

import (
	"fmt"
	"math"
	"time"
)

func add(x, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}

var f, k, j bool

func main() {

	hola, ja, sapo := "ahol", true, 2
	fmt.Println("hola mundo")
	fmt.Println(time.Now())
	fmt.Println("tienees %g problemas.\n", math.Sqrt(16))
	fmt.Println(add(3, 2))

	a, b := swap("hello", "world")
	fmt.Println(b, a)
	var i int
	fmt.Println(f, k, j, i)
	fmt.Println(hola, ja, sapo)

}
