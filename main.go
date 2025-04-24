package main 

import "fmt"      
func pegarnome() (string, string) {
	return "Barry", "Allen"
}

func main() {
nome, sobrenome := pegarnome()
fmt.Println(nome)
fmt.Println(sobrenome)
}