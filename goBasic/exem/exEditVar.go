package main

import "fmt"

func main() {
	nome := "Douglas"
	idade := 24
	versao := 1.1
	fmt.Println("Olá sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scan(&comando)
	// Valor da variavel na memoria
	fmt.Print("valor do resultado na memerio:", &comando)
	// Valor da variavel digitado
	fmt.Println("\nO comando escolhido foi:", comando)
}
