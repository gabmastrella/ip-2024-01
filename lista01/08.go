package main

import "fmt"

var r, a, p float32

func main() {

	fmt.Print("Insira o raio da lata, em metros: ")
	fmt.Scan(&r)

	fmt.Print("Insira a altura da lata, em metros: ")
	fmt.Scan(&a)

	p = 3.14159
	Al := 2 * p * r * a
	Ac := p * (r * r)
	At := 2*Ac + Al
	X := At * 100

	fmt.Printf("O VALOR DO CUSTO E = %.2f\n", X)

}
