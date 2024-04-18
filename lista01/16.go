package main

import "fmt"

var s float32

func main() {

	fmt.Print("Insira o valor do salário: ")
	fmt.Scan(&s)

	if s <= 300 {

		x := s + (0.5 * s)
		fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", x)

	} else {

		x := s + (0.3 * s)
		fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", x)

	}

}
