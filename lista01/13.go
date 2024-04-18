package main

import "fmt"

var x float32
var c int = 0

func main() {

	fmt.Print("Insira sua nota: ")
	fmt.Scan(&x)

	for x >= 0 && x < 6 && c < 1 {
		y := "D"
		fmt.Printf("NOTA = %.1f CONCEITO = %s\n", x, y)
		c = c + 1
	}
	for x >= 6 && x < 7.5 && c < 1 {
		y := "C"
		fmt.Printf("NOTA = %.1f CONCEITO = %s\n", x, y)
		c = c + 1
	}
	for x >= 7.5 && x < 9 && c < 1 {
		y := "B"
		fmt.Printf("NOTA = %.1f CONCEITO = %s\n", x, y)
		c = c + 1
	}
	for x >= 9 && x < 10 && c < 1 {
		y := "A"
		fmt.Printf("NOTA = %.1f CONCEITO = %s\n", x, y)
		c = c + 1
	}

}
