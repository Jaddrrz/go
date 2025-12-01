package main

import "fmt"
import "io"

func main() {
  var parola string
  var lista []string
  var media int
  somma, c := 0,0

  for {
    _,err :=fmt.Scan(&parola)
    if err == io.EOF {
      break
    }
    lista = append(lista, parola)
    somma += len(parola)
    c++
  }
  media = int(float64(somma)/float64(c))

  for _,v := range lista {
    if len(v) > media {
      fmt.Printf("%s ",v)
    }
  }
  fmt.Println()
}
