package main

import "fmt"

func merge(x []int, y []int) []int {
  var z []int
  var i1, i2, k int
  
  for i1 < len(x) && i2 < len(y) {
    if x[i1] == y[i2] {
        //i1 += 1
      // i2 += 1
      k += 1
      i1 += 1
      continue
    }
    if x[i1] < y[i2] {
      z = append(z, x[i1])
      i1 += 1
    } else {
      z = append(z, y[i2])
      i2 += 1
    }
    k += 1
  }
  if i1 < len(x) {
    z = append(z, x[i1:]...)
  } else {
    z = append(z, y[i2:]...)
  }
  return z
}

func main() {
  list1 := []int{1,2,4,7,10}
  list2 := []int{3,4,6,10}

  fmt.Println(list1)
  fmt.Println(list2)
  fmt.Println(merge(list1, list2))
}
