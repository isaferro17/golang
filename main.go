package main 
import "fmt"

func main() {
// Slice
var numeros = []int {2,4,8,10,12}
fmt.Println(numeros)
numeros[3] = 5
fmt.Println(numeros, len(numeros), cap(numeros))
numeros = append(numeros,14, 16, 18 )
fmt.Println(numeros, len(numeros), cap(numeros))
}