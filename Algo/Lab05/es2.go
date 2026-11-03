package main

import "fmt"
import "math/rand"
import "os"
import "strconv"

func bubbleSort(s []int) []int {
  for i := 0; i < len(s); i++ {
    for j := 0; j < len(s)-1; j++ {
      if s[j] > s[j+1] {
        s[j], s[j+1] = s[j+1], s[j]
      } 
    }
  }
  return s
}

func merge(x, y []int) []int {
  var z []int
  var i1, i2 int
  
  for i1 < len(x) && i2 < len(y) {
    if x[i1] < y[i2] {
      z = append(z, x[i1])
      i1 += 1
    } else {
      z = append(z, y[i2])
      i2 += 1
    }
    //k += 1
  }
  if i1 < len(x) {
    z = append(z, x[i1:]...)
  } else {
    z = append(z, y[i2:]...)
  }
  return z
}

func mergeSort(x []int) []int {
  if len(x) > 1 {
    var b, c []int
    m := len(x) / 2
    b = append(b, x[:m]...)
    c = append(c, x[m:]...)
    b = mergeSort(b)
    c = mergeSort(c)
    x = merge(b, c)
  }
  return x
}

func generateArray(n int) []int {
  var res []int
  for i := 0; i < n; i++ {
    res = append(res, 1 + rand.Intn(9))
  }
  return res
}

func fillArray(list [][]int, k, n int) {
  for i := 0; i < k; i++ {
    list[i] = append(list[i], generateArray(n)...)
  }
}

func mergeMultiple(list [][]int) []int {
  if len(list) == 2 {
    return merge(bubbleSort(list[0]), bubbleSort(list[1]))
  } else {
    return merge(list[0], mergeMultiple(list[1:]))
  }
}

func main() {
  dim, _ := strconv.Atoi(os.Args[1])
  n, _ := strconv.Atoi(os.Args[2])
  list := make([][]int, dim)

  
  fillArray(list, dim, n)
  fmt.Println(mergeSort(list[0]))
  fmt.Println(list)
  fmt.Println(mergeMultiple(list))

}

