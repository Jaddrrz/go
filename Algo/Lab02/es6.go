package main

import "fmt"

func nZeri(n int) int {
  if n == 0 {
    return 0
  } else {
    if n % 10 == 0 {
      return 1 + nZeri(n/10)
    } 
  }
  return 0 + nZeri(n/10)
}

func main() {
  var n int
  fmt.Scan(&n)
  fmt.Println(nZeri(n))
}
