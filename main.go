package main 
import "fmt"

func main() {
var usuario, senha string
fmt.Println("Digite corretamente o usuario")
fmt.Scan(&usuario)

fmt.Println("Digite a senha")
fmt.Scan(&senha)


if usuario== "admin" && senha== "1234" {
fmt.Println("Bem vindo")
} else {
fmt.Println("usuario e senha incorretos")
}

}