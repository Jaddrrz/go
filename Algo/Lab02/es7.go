package main

import "fmt"
//import "math"
import "math/big"

func potenza(n int) int {
  res := 1
  for i := 0; i < n; i++ {
    res *= n
  }
  return res
}

func main() {
  var n int = 
  //fmt.Scan(&n)
  fmt.Println(potenza(n))
}
