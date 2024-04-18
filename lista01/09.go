package main

import "fmt"

var A, B, C, B2 float32

func main() {

	fmt.Print("INSIRA O VALOR DE A: ")
	fmt.Scan(&A)
	fmt.Print("INSIRA O VALOR DE B: ")
	fmt.Scan(&B)
	fmt.Print("INSIRA O VALOR DE C: ")
	fmt.Scan(&C)

	X := (B * B) - (4 * A * C)

	fmt.Printf("O VALOR DE DELTA E = %.2f\n", X)

}
