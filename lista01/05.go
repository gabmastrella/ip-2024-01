package main

import "fmt"

var u int
var cons float32
var tipo string
var controle int = 0

func main() {

	fmt.Println("INSIRA A CONTA, CONSUMO DE ÁGUA POR M³ E O TIPO DE CONSUMIDOR:")
	fmt.Scanf("%d %f %v", &u, &cons, &tipo)

	for (tipo == "R") && (controle < 1) {
		v := 5 + (0.05 * cons)
		fmt.Printf("CONTA = %d\n", u)
		fmt.Printf("VALOR DA CONTA = %.2f\n", v)
		controle = controle + 1
	}
	for (tipo == "C") && (controle < 1) {
		v := 500 + (0.25 * (cons - 80))
		fmt.Printf("CONTA = %d\n", u)
		fmt.Printf("VALOR DA CONTA = %.2f\n", v)
		controle = controle + 1
	}
	for (tipo == "I") && (controle < 1) {
		v := 800 + (0.04 * (cons - 100))
		fmt.Printf("CONTA = %d\n", u)
		fmt.Printf("VALOR DA CONTA = %.2f\n", v)
		controle = controle + 1
	}

}
