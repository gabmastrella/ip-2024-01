package main

import "fmt"

var h, m, s int

func main() {

	fmt.Scan(&h)
	fmt.Scan(&m)
	fmt.Scan(&s)

	X := (h * 3600) + (m * 60) + s
	fmt.Printf("O TEMPO EM SEGUNDOS E = %d\n", X)

}
