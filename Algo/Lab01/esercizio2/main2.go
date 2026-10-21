package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "time"

func main() {
  start := time.Now()
  var res []string
  nomeFile := os.Args[1]
  file, _ := os.Open(nomeFile)
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    res = append(res, line)
  }
  fmt.Println(strings.Join(res, " "))
  timeElapsed := time.Since(start)
  fmt.Printf("Tempo di esecuzione %v\n", timeElapsed)
}
