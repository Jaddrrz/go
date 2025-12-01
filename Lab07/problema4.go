package main

import "fmt"

func main() {
  var parola string
  var pari, dispari []string

  for {
    fmt.Scan(&parola)
    if parola == "stop" {
      break
    }
    if len(parola) % 2 == 0 {
      pari = append(pari, parola)
    } else {
      dispari = append(dispari, parola)
    }
  }
  
  for _,v := range pari {
    fmt.Printf("%s ", v)
  }
  fmt.Println()
  for _,v := range dispari {
    fmt.Printf("%s ", v)
  }
  fmt.Println()
}
