package main 

import "fmt"      
func  dadospessoais() (name string,idade int) {

fmt.Println("digite seu nome")
fmt.Scan(&name)
fmt.Println("Digite sua idade")
fmt.Scan(&idade)
return name, idade
}

func main(){
var name, idade = dadospessoais()	
if idade < 18 {
 fmt.Println(name + ". Você é menor de idade")

 } else{
	fmt.Println(name + ". Você é maior de idade")
}
}