package main

import "fmt"
import "math/rand"
import "os"
import "strconv"

func horner(polinomio []int, x int, i int) int {
  if i == len(polinomio) - 1 {
    return polinomio[i]
  } else {
    return polinomio[i] + x * horner(polinomio, x, i + 1)
  }
}

func generaPolinomio() []int {
  n := 3
  var res []int
  for i := 0; i < n; i++ {
    res = append(res, 1+ rand.Intn(10))
  }
  return res
}

func instertionHorner(polinomi [][]int) [][]int {
  for i := 1; i < len(polinomi); i++ {
    x := polinomi[i]
    j := i - 1
    for l := 1; l <= 1000; l++ {
      sum1 += horner(polinomi[0], l, 0)
      sum2 += horner(polinomi[1], l, 0)
    }
    for j >= 0 && horner(polinomi[j]) > horner(x) {
      polinomi[j + 1] = polinomi[j]
      j -= 1
    }
    polinomi[j + 1] = x
  }
  return polinomi
}

func main() {
  var polinomi [][]int
  var sum1, sum2 int
  //list := []int{1,10,23,4,24,17}
  n, _ := strconv.Atoi(os.Args[1])
  for i := 0; i < n; i++ {
    polinomi = append(polinomi, generaPolinomio())
  }
  
  fmt.Println(polinomi)
  fmt.Println(sum1, sum2)
  fmt.Println(instertionHorner(polinomi))
}
