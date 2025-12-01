package main

import "fmt"

func main() {
  var parola string

  for {
    fmt.Scan(&parola)
    if parola == "stop" {
      break
    }
    chars := []rune(parola)
    var nuovaParola string
    for _,v := range chars {
      if (v >= 97 && v <= 122) || (v >= 65 && v <= 90) {
          nuovaParola += string(v)
      }
    }
    fmt.Printf("%s ", nuovaParola)
  }
  fmt.Println()
}
