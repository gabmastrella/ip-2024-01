package main

import "fmt"

var x, y int

func main() {

	fmt.Print("Insira os valores: ")
	fmt.Scan(&x, &y)

	if x%2 == 0 {

		for i := 0; i < y; i++ {

			fmt.Printf("%d ", x)

			x = x + 2
		}

	} else {

		fmt.Print("O PRIMEIRO NUMERO NAO E PAR\n")

	}

}
