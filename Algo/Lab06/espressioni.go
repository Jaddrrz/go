package main

import "bufio"
import "fmt"
import "os"
//import "github.com/emirpasic/gods/stacks/linkedliststack"

// Controlla se il numero di parentesi è bilanciato
func isBalanced(s string) bool {
  var open, close int
  
  for _, v := range s {
    if string(v) == "(" {
      open++
    } else if string(v) == ")" {
      close++
    }
  }
  
  if open != close {
    return false
  }
  return true
}

// Controlla se l'espressione è valida
func isValid(s string) bool {
  return true
}

// A ogni coppia di parentesi crea una pila, considerare pila simboli e pila cifre

func main() {
  var input string
  scanner := bufio.NewScanner(os.Stdin)

  if scanner.Scan() {
    input = scanner.Text()
  }
  fmt.Println(isBalanced(input))
}
