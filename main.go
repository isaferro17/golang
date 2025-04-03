package main 
import "fmt"

func main() {
var ages = [4] int {16,17,10,40}
nomes := [4] string{"Mario, Luigi, Red hood, Constantine"}
fmt.Println(ages)
fmt.Println(nomes)
// Slice
var score = []int {100,200,300,400}
fmt.Println(score)
score[1] = 2
fmt.Println(score, len(score), cap(score))
rangeOne := score [1:3]
fmt.Println(rangeOne)
rangeTwo:= score [2:]
fmt.Println(rangeTwo)
rangeThree:= score[:3]
fmt.Println(rangeThree)
var superherois = []string{"homem-aranha,batman, greenlantern"} 
fmt.Println(superherois)
superherois = append(superherois,"Ben 10" )
fmt.Println(superherois, len(superherois), cap(superherois))

}