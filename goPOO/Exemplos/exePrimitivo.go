package main

import "fmt"

// struct = se torna um grande varaivel
type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	var titular string = "Guilherme"
	var numeroAgencia int = 589
	var numeroConta int = 123456
	var saldo float64 = 125.50

	var titular2 string = "Luciene"
	var numeroAgencia2 int = 555
	var numeroConta2 int = 111333
	var saldo2 float64 = 200

	fmt.Println(titular, numeroAgencia, numeroConta, saldo)
	fmt.Println(titular2, numeroAgencia2, numeroConta2, saldo2)
}
