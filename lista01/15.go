package main

import "fmt"

var N int
var controle int = 2
var saida string

func main() {

	fmt.Scan(&N)

	for 5 < N && N < 2000 && controle <= N {

		a := controle
		b := a * a

		controle = controle + 2

		x := fmt.Sprintf("%d", a)
		y := fmt.Sprintf("%d", b)

		saida += string(x) + "^2 = " + string(y) + "\n"

	}
	fmt.Print(saida)
}
