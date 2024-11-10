package main

import (
	"fmt"
)

func main() {
	// DECLARAÇÃO DAS VARIÁVEIS
	var altura, base, area float32

	// ENTRADA DE DADOS
	fmt.Print("Valor da base do triângulo: ")
	fmt.Scan(&base)
	fmt.Print("Valor da altura do triângulo: ")
	fmt.Scan(&altura)

	// CÁLCULO DA ÁREA
	area = (base * altura) / 2

	// SAÍDA
	fmt.Printf("A ÁREA DO TRIÂNGULO É: %v", area)
}
