package main 

import "fmt"      

var saldo float64

func depositar(valor float64) {
	saldo += valor 
	fmt.Println("Seu saldo atual é:", saldo)
}

func sacar(valor float64) {
	if valor > saldo {
		fmt.Println("Saldo insuficiente para saque.")
	} else {
		saldo -= valor
		fmt.Println("Saque realizado. Seu saldo atual é:", saldo)
	}
}

func main(){
	saldo = 200
	var opção int
	fmt.Println("Digite 1 para depositar ou 2 para sacar:")
	fmt.Scan(&opção)
	if opção == 1  {
		var valor float64
		fmt.Println("Digite o valor")
		fmt.Scan(&valor)
		depositar(valor)
	} else {
		var valor float64
		fmt.Println("Digite o valor")
		fmt.Scan(&valor)
		sacar(valor)
	}
}

