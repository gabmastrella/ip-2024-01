package main

import "fmt"

var x, y, n, saida int

func main() {

	fmt.Print("Insira os valores: ")
	fmt.Scan(&x, &y, &n)

	for i := 0; i < n; i++ {

		saida += x + (i * y)

	}

	fmt.Printf("%d\n", saida)

}
