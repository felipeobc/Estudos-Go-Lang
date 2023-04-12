package main

import "fmt"

func main() {
	var precoLeite float32 = 2.99
	var precoOvo float32 = 3.99
	var precoPao float32 = 0.99
	var nome string = "Felipe" // nao tem a nescessidade de usar tipos de variaveis
	var versao float32 = 1.1
	var idade int = 31

	fmt.Println("Hello Sir", nome, "Sua idade é:", idade)
	fmt.Println("Esse programa está: ", versao)

	fmt.Println("O preço do leite é R$", precoLeite)
	fmt.Println("O preço do ovo é R$", precoOvo)
	fmt.Println("O preço do pão é R$", precoPao)
}
