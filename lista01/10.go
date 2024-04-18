package main

import "fmt"

var a, b, c, d float32

func main() {

	fmt.Scanf("%f\n", &a)
	fmt.Scanf("%f\n", &b)
	fmt.Scanf("%f\n", &c)
	fmt.Scanf("%f\n", &d)

	X := (a * d) - (b * c)

	fmt.Printf("O VALOR DO DETERMINANTE E = %.2f\n", X)
}
