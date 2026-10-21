package main

import "fmt"
import "strconv"
import "os"

func occorrenze(s string, char string, i int) int {
  if len(s) == 0 {
    return 1
  } else {
    if string(s[i]) == char {
      i += 1
      fmt.Println("+1")
      return 1 + occorrenze(s[i:len(s)-i], char, i)
    }
  }
  return 0
}

func main() {
  stringa := os.Args[1]
  char := os.Args[2]
  i,_ := strconv.Atoi(os.Args[3])
  fmt.Println(occorrenze(stringa, char, i))
}
