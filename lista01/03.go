package main

import "fmt"

var a, b, c int

func main() {

	fmt.Println("Informe 3 valores menores que 10: ")
	fmt.Scan(&a, &b, &c)

	if a >= 10 {
		fmt.Println("DIGITO INVALIDO")
	} else {
		if b >= 10 {
			fmt.Println("DIGITO INVALIDO")
		} else {
			if c >= 10 {
				fmt.Println("DIGITO INVALIDO")
			} else {

				conc := a*100 + b*10 + c
				qd := (a*100 + b*10 + c) * (a*100 + b*10 + c)

				fmt.Printf("%d, %d\n", conc, qd)

			}
		}
	}
}
