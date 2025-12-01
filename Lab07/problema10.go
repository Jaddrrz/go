package main

import "fmt"
import "io"

func main() {
  var parola string
  var max int
  var lista []string

  for {
    _,err := fmt.Scan(&parola)
    if err == io.EOF {
      break 
    }
    lista = append(lista, parola)
  }

  max = len(lista[0])
  for i := 0; i < len(lista)-1; i++ {
    if len(lista[i]) > max {
      max = len(lista[i])
    }
  }

  for i := 0; i < len(lista)-1; i++ {
    if len(lista[i]) == max {
      fmt.Println(lista[i])
    }
  }
  
}
