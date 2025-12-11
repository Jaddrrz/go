package main

import "bufio"
import "fmt"
import "strconv"
import "strings"
import "os"

type Segmento struct {
  estremi, interno byte
  orizzontale bool
  lunghezza int
}

func (s Segmento) String() {
  fmt.Print(string(s.estremi))
  if !s.orizzontale {
    fmt.Println()
  }
  for i := 0; i < s.lunghezza - 2; i++ {
    if s.orizzontale {
      fmt.Print(string(s.interno))
    } else {
      fmt.Println(string(s.interno))
    }
  }
  fmt.Print(string(s.estremi))
  fmt.Println()
}

func allunga(s *Segmento, n int) {
  s.lunghezza += n
}

func main() {
  var s Segmento
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  input := strings.Fields(scanner.Text())
  //estremi := []byte(input[0])
  s.estremi = []byte(input[0])[0]
  //interno := []byte(input[1])
  s.interno = []byte(input[1])[0]
  s.orizzontale, _ = strconv.ParseBool(input[2])
  s.lunghezza, _ = strconv.Atoi(input[3])
  fmt.Println(s)
  s.String()
  allunga(&s, 3)
  s.String()
}

