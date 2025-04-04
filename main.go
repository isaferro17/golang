package main 

import (
  "fmt"      
  "sort"
  "strings"
  )

func main() {
  
  greeting := "hello my friends"
  fmt.Println(strings.Contains(greeting,"dogs"))
  fmt.Println(strings.ReplaceAll(greeting,"my", "mine"))
  fmt.Println(strings.ToUpper(greeting))
  fmt.Println(strings.Index(greeting,"my"))
  fmt.Println(strings.Split(greeting,""))
  ages:= []int{50,80,10}
  sort.Ints(ages)
  fmt.Println(ages)
  index:= sort.SearchInts(ages,50)
fmt.Println(index)
names := []string{"Alice", "marco", "Diego"}
sort.Strings(names)
fmt.Println(names)
fmt.Println(sort.SearchStrings(names,"Alice"))
x:=0
for x < 5 {
  fmt.Println(x)
  x++
  }
  for i:=0; i <5; i++{
    fmt.Println("for 2:")
  }
  for i:=0; i <len(names); i++{
    fmt.Println(names)
}

for index, value := range names { 
  fmt.Println("o indice é", index, "e o valor", value)
}
for index, value := range ages{
fmt.Println("o índice é", index, "e o valor", value)
}

}