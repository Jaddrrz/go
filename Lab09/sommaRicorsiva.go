package main

import "bufio"
import "fmt"
import "strconv"
import "os"

func recursiveSum(list []int) int {
  if len(list) == 0 {
    return 0
  } else {
    return list[0] + recursiveSum(list[1:])
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
  fmt.Println(recursiveSum(list))
}
