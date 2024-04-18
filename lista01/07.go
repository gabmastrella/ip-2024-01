package main

import "fmt"

var f, c float32

func main() {

	fmt.Print("Insira a temperatura em Fahrenheit: ")
	fmt.Scan(&f)

	fmt.Print("Insira a quantidade de chuva em polegadas: ")
	fmt.Scan(&c)

	X := (5 * (f - 32)) / 9
	Y := c * 25.4

	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", X)
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f\n", Y)
}
