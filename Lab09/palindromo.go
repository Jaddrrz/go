package main

import "fmt"

func isPalindrome(s string) bool {
  if len(s) == 1 || len(s) == 0 {
    return true
  } else {
    if s[0] == s[len(s)-1] {
      return isPalindrome(s[1:len(s)-1])
    } else {
      return false
    }
  }
}

func main() {
  var s string
  fmt.Scan(&s)
  fmt.Println(isPalindrome(s))
}
