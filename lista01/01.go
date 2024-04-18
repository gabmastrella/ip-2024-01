package main

import "fmt"

var a float32
var b float32
var c float32

func main() {
	fmt.Print("Insira os valores: ")
	fmt.Scanf("%f%f%f", &a, &b, &c)

	q := a + b + c
	media := q / 3

	fmt.Printf("MEDIA = %.2f\n", media)

	if media < 6 {
		fmt.Println("REPROVADO")
	} else {
		fmt.Println("APROVADO")
	}

}
