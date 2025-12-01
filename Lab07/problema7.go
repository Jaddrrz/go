package main

import "fmt"

func main() {
  var parola string
  var lista []string

  for {
    fmt.Scan(&parola)
    if parola == "stop" {
      break
    }
    lista = append(lista, parola)
  }

  for i := 0; i < len(lista)-1; i++ {
    for j := 0; j < len(lista)-i-1; j++ {
      if lista[j] > lista[j+1] {
        lista[j], lista[j+1] = lista[j+1], lista[j]
      }
    }
  }
  for _,v := range lista {
    fmt.Printf("%s ",v)
  }
  fmt.Println()
}
