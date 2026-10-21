package main

import "fmt"
//import "strconv"
//import "os"

func isOrdinato(list []int) bool {
  if len(list) == 1 {
    return true
  } else {
    if list[0] < list[1] {
      return isOrdinato(list[1:])
    } 
    return false
  }
}

func main() {
  input1 := []int{10,11,1,4,5,7,20}
  input2 := []int{2,4,7}

  fmt.Println(isOrdinato(input1))
  fmt.Println(isOrdinato(input2))
}
