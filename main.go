package main 

import "fmt"      

func main () {
 estoque := map[string]int{}
 estoque["Coxinha"] = 10
 estoque["Pão de Queijo"] = 15
 estoque["Refrigerante"] = 20
estoque["Coxinha"] -=2
estoque["Pão de Queijo"] -=1
 

fmt.Println("Estoque Atual:")
for produto, quantidade := range estoque {
	fmt.Printf("%s: %d\n", produto, quantidade)
}

}