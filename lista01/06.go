package main

import "fmt"

var n int
var controle int = 0
var x float32
var saida string

func main() {

	fmt.Print("Insira a quantidade de conversões a serem realizadas: ")
	fmt.Scanf("%d\n", &n)

	for controle < n {

		fmt.Printf("Informe a temperatura: ")
		fmt.Scanf("%f\n", &x)

		controle = controle + 1

		y := (5 * (x - 32)) / 9

		ystring := fmt.Sprintf("%.2f", y)
		xstring := fmt.Sprintf("%.2f", x)

		saida += (xstring + " FAHRENHHEIT EQUIVALE A " + ystring + " CELSIUS\n")

	}

	if controle == n {
		fmt.Println(saida)
	}
}
