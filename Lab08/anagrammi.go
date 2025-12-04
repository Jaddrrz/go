package main

import "fmt"
import "os"

func sortSlice(s []rune) {  
  for i := 0; i < len(s)-1; i++ {
    for j := 0; j < len(s)-i-1; j++ {      
      if s[j] > s[j+1] {
        s[j], s[j+1] = s[j+1], s[j]
      }
    }
  }
}

func isAnagram(s1, s2 string) bool {
  var elemsAreEqual bool
  m1, m2 := []rune(s1), []rune(s2)
  sortSlice(m1)
  sortSlice(m2)
  for i := 0; i < len(m1); i++ {
    if m1[i] == m2[i] {
      elemsAreEqual = true
    } else {
      return false
    }
  }
  return elemsAreEqual
}

func main() {
  if len(os.Args) < 3 || len(os.Args) > 3 {
    fmt.Println("input errato")
    return
  }
  p1, p2 := os.Args[1], os.Args[2]
  if isAnagram(p1,p2) {
    fmt.Println("Le due stringhe sono una l'anagramma dell'altra.")
  } else {
    fmt.Println("Le due stringhe non sono una l'anagramma dell'altra.")
  }
}
