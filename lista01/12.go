package main

import (
	"fmt"
	"math"
)

var horas, t float64

func main() {

	fmt.Scan(&horas)

	t = math.Mod(horas, 3)

	X := ((horas - t) / 3 * 10) + (t * 5)

	fmt.Printf("O VALOR A PAGAR E = %.2f\n", X)
}
