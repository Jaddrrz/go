package main

import "fmt"
import "math/rand"

func sort(s []int) []int {
  for i := 0; i < len(s); i++ {
    for j := 0; j < len(s)-1; j++ {
      if s[j] > s[j+1] {
        s[j], s[j+1] = s[j+1], s[j]
      } 
    }
  }
  return s
}

func ricercaBin(s []int, n int) bool {
  m := len(s)/2
  if len(s) == 0 {
    return false
  }
  if s[m] == n {
    return true
  } else if n < s[m] {
    return ricercaBin(s[:m], n)
  } else {
    return ricercaBin(s[m+1:], n)
  }
  return false
}

func main() {
  var lista []int
  var p, q int
  fmt.Scan(&p, &q)

  for i := 0; i < p; i++ {
    lista = append(lista, rand.Intn(10))
  }
  fmt.Println(lista)
  if ricercaBin(sort(lista), q) {
    fmt.Println("trovato")
  } else {
    fmt.Println("non trovato")  
  }
}
