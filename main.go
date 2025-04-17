package main 

import "fmt"      
func dividir (dividendo int , divisor int) (int,string) {
	if divisor == 0 {
		return 0, "Erro na divisao por zero"
}
return dividendo / divisor , "sem erro"
}
func operacoesbasicas(a int, b int) (int, int,int) {
	soma := a + b
	multiplicacao := a * b
	subtracao := a - b
return soma, multiplicacao, subtracao
}

func main(){
resultado, erro := dividir (10,0)
if erro != "sem erro" {
fmt.Println(erro)
} else {
	fmt.Println("O resultado da divisão é",resultado, erro)

}
soma, multi, sub := operacoesbasicas(10,2)
fmt.Println(soma)
fmt.Println(multi)
fmt.Println(sub)
}