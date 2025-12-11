package main

import "fmt"

func fattoriale(n int) int {
  if n == 0 {
    return 1
  } else {
    return n * fattoriale(n-1)
  }
}

func main() {
  var n int
  fmt.Scan(&n)
  fmt.Println(fattoriale(n))
}
