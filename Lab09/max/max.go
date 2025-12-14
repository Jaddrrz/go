package main

import "bufio"
import "fmt"
import "os"

import "strconv"

func recursiveMax(list []int) int {
  fmt.Println(list[0])
  if len(list) == 1 {
    return list[0]
  } else {
    for i := 0; i < len(list); i++ {
      if list[i] > list[0] {
        list[0] = list[i]
      }
    }
    return recursiveMax(list[:len(list)-1])
  }
}

func main() {
  var list []int
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    if scanner.Err() != nil {
      break
    }
    riga := scanner.Text()
    n, _ := strconv.Atoi(riga)
    list = append(list, n)
  }
  fmt.Println(recursiveMax(list))
}
