package main

import "bufio"
import "fmt"
import "os"
import "strings"

func contaVocali(s string, vocali map[rune]int) {
    for _,v := range s {
      if strings.ContainsRune("AEIOUaeiou", v) {
        vocali[v] += 1
      }
    }
}

func main() {
  vocali := make(map[rune]int)
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    if err := scanner.Err(); err != nil {
      break
    }
    riga := scanner.Text()
    contaVocali(riga, vocali)
  }
  for i,v := range vocali {
    fmt.Printf("%s : %d\n", string(i), v)
  }
}
