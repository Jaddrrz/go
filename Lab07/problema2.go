package main

import "fmt"

func main() {
  var parola string
  for {
    fmt.Scan(&parola)
    if parola == "stop" {
      break
    }
    fmt.Println(parola)
  }
}
