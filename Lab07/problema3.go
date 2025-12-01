package main

import "fmt"

func main() {
  var parola string
  for {
    fmt.Scan(&parola)
    if parola == "stop" {
      break
    }
    if len(parola) % 2 == 0 {
      fmt.Println(parola)
    }
  }
}
