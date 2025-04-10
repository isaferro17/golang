package main 

import "fmt"      

func main() {
	age:= 45
  fmt.Println(age < 50)
  fmt.Println(age > 50)
  fmt.Println(age == 50)
  fmt.Println(age != 50)

  if age < 30 {
    fmt.Println("Menor que 30 anos")
  } else if age < 40 {
    fmt.Println("menor que 40 anos") 
  }else{
    fmt.Println("menor que 40 anos")
  }
  nomes := []string {"Davos","Kevan","Myrcella","Alyssanne","Jace", "Luke"}

  for index, value := range nomes {
    if index == 1 {
      fmt.Println("continue após a posição", index,"e valor", value)
      continue
    }
    if index > 2 {
    fmt.Println("Sair após",index)
    break
  }
  fmt.Println("valor:", value)
  }
}