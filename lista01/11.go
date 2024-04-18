package main

import "fmt"

var n int

func main() {

	fmt.Print("Insira um valor: ")
	fmt.Scan(&n)

	if n%15 == 0 {
		fmt.Print("O NUMERO E DIVISIVEL\n")
	} else {
		fmt.Print("O NUMERO NAO E DIVISIVEL\n")
	}

}
