package main

import "fmt"

var x int
var saida float32

func main() {

	fmt.Print("Insira o valor: ")
	fmt.Scan(&x)

	if x <= 1 {

		fmt.Print("Numero Invalido!")

	} else {

		for i := 0; i < x; i++ {

			saida += 1 / (float32(i) + 1)

		}
		fmt.Printf("%f\n", saida)
	}
}
