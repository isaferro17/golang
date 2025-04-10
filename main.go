package main 

import "fmt"      

func main() {
	var s, v float64
	var c string
	fmt.Print("Saldo inicial: ")
	fmt.Scan(&s)
	for {
		fmt.Print("Comando (sacar/depositar/parar): ")
		fmt.Scan(&c)
		if c == "sacar" {
			fmt.Print("Valor: ")
			fmt.Scan(&v)
			s -= v
		} else if c == "depositar" {
			fmt.Print("Valor: ")
			fmt.Scan(&v)
			s += v
		} else if c == "parar" {
			break
		}
		fmt.Println("Saldo:", s)
	}
}
