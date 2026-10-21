package main

import "fmt"
import "os"

func palindromo(s string) bool {
  if len(s) == 0 || len(s) == 1 {
    return true
  } else {
    if s[0] == s[len(s)-1] {
      return palindromo(s[1:len(s)-1])
    }
    return false
  }
}

func main() {
  input := os.Args[1]
  fmt.Println(palindromo(input))
}
