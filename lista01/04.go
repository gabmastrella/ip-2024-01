package main

import "fmt"

//100kW = 70% do salário mínimo
var sM, kW, kWg float32

func main() {

	fmt.Print("Insira o salário mínimo: ")
	fmt.Scanf("%f \n", &sM)

	kW = 0.7 * sM / 100

	fmt.Print("Insira a quantidade de kW gasta: ")
	fmt.Scanf("%f \n", &kWg)

	fmt.Printf("\nCusto por kW: R$ %.2f\n", kW)

	ckW := kW * kWg

	fmt.Printf("Custo do Consumo: R$ %.2f\n", ckW)

	ckWd := ckW * 0.9

	fmt.Printf("Custo do Consumo com Desconto: R$ %.2f\n", ckWd)
}
