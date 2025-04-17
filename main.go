package main 

import "fmt"      
func main(){
alunoidade := make(map[string]int)
alunoidade ["Bruno"] = 15
alunoidade ["Otávio"] = 16 
alunoidade ["fabiano"] = 40
alunoidade ["Isabela"] = 15
fmt.Println("idade da isabela", alunoidade,["Isabela"])

notasalunos: map [string]float64{
"Bruno" : 9,7,
"Otávio" : 10,
"fabiano" : 8,7,
"Isabela" : 9,7,
}
for k, v := range notasalunos{
fmt.Println("%s tirou a nota mais %f /n", k,v)
}
}