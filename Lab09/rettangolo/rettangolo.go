package main

import "fmt"

type Rettangolo struct {
  base, altezza int 
}

func (r Rettangolo) String() {
  for x := 1; x <= r.base; x++ {
    for y := 1; y <= r.altezza; y++ {
      fmt.Print(".")
    }
    fmt.Println()
  }

  
}

func main() {
  var r Rettangolo
  fmt.Scan(&r.base, &r.altezza)
  r.String()
}
