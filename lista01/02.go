package main

import "fmt"

var controle int = 0
var teste int
var qtotal float32
var perc_pop float32
var perc_geral float32
var perc_arq float32
var perc_cad float32
var saida string

func main() {

	fmt.Print("Informe a quantidade de jogos: ")
	fmt.Scan(&teste)

	for controle < teste {
		fmt.Print("Informe os valores: ")
		fmt.Scan(&qtotal, &perc_pop, &perc_geral, &perc_arq, &perc_cad)
		controle = controle + 1

		tot_pop := qtotal * (perc_pop / 100)
		tot_geral := qtotal * (perc_geral / 100) * 5
		tot_arq := qtotal * (perc_arq / 100) * 10
		tot_cad := qtotal * (perc_cad / 100) * 20

		rendaTotal := tot_arq + tot_cad + tot_geral + tot_pop
		rstring := fmt.Sprintf("%.2f", rendaTotal)
		cstring := fmt.Sprintf("%v", controle)
		saida += ("A RENDA TOTAL DO JOGO N. " + string(cstring) + " E DE: " + string(rstring) + "\n")

	}

	if controle == teste {
		fmt.Println(saida)
	}

}
