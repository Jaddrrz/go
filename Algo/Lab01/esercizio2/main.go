package main

import "fmt"
import "os"
import "bufio"
import "strings"

func main() {
  var res string
  nomeFile := os.Args[1]
  file, _ := os.Open(nomeFile)
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    char := strings.Split(line, "")[0]
    res = res + char + " "
  }
  fmt.Println(res)
}
