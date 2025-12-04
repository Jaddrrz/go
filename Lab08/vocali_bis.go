package main

import "bufio"
import "fmt"
import "os"
//import "sort"
import "strings"

func contaVocali(s string, vocali map[rune]int) {
    for _,v := range s {
      if strings.ContainsRune("AEIOUaeiou", v) {
        vocali[v] += 1
      }
    }
}

func main() {
  var vocali = []rune("AEIOUaeiou")
  res := make(map[rune]int)
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    if err := scanner.Err(); err != nil {
      break
    }
    riga := scanner.Text()
    contaVocali(riga, res)
  }
  
  for _,v := range vocali {
    if res[v] == 0 {
      continue
    }
    fmt.Printf("%s : %d\n", string(v), res[v])
  }
}
