package main

import "fmt"

func reverseString(s string) string {
  sl := []rune(s)
  if len(sl) == 0 {
    return ""
  } else {
    
  }
}

func main() {
  var s string
  fmt.Scan(&s)
  fmt.Println(reverseString(s))
}
