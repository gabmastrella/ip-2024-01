package main

import (
	"fmt"
	"math"
)

var h, a float64

func main() {

	fmt.Scan(&h)
	fmt.Scan(&a)

	b := math.Sqrt(3)
	Ab := (3 * (a * a) * b) / 2
	X := (Ab * h) / 3

	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS\n", X)

}
