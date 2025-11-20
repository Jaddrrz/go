package main

import "fmt"
import "math/rand"
import "os"
import "os/exec"
import "strconv"
import "time"

type Mosca struct {
  x, y int
}

func stampaStanza(n int, m Mosca) {
  for x := 1; x <= n; x++ {
    for y := 1; y <= n; y++ {
      switch {
      case x == m.x && y == m.y:
        fmt.Print("*")
      case x == 1 && y == 1:
        fmt.Print("\u250C")
      case x == 1 && y == n:
        fmt.Print("\u2510")
      case x == n && y == 1:
        fmt.Print("\u2514")
      case x == n && y == n:
        fmt.Print("\u2518")
      case (x > 1 && x < n) && (y == 1 || y == n):
        fmt.Print("\u2502")
      case (y > 1 && y < n) && (x == 1 || x == n):
        fmt.Print("\u2500")
      default:
        fmt.Print(" ")
      }
    }
    fmt.Println()
  }
}

func spostaMosca(dimStanza int, m *Mosca) {
  n := rand.Intn(4)
  switch n {
  case 0:
    m.x++
    if m.x == dimStanza {
      m.x--
    }
  case 1:
    m.x--
    if m.x == 1 {
      m.x++
    }
  case 2:
    m.y++
    if m.y == dimStanza {
      m.y--
    }
  case 3:
    m.y--
    if m.y == 1 {
      m.y++
    }
  }
}

func cancella() {
   cmd := exec.Command("clear")
   cmd.Stdout= os.Stdout
   cmd.Run()
}


func main() {
  var n int
  var mosca Mosca
  
  n, _ = strconv.Atoi(os.Args[1])
  mosca = Mosca{rand.Intn(n-2) + 2, rand.Intn(n-2) + 2}  

  cancella()
  stampaStanza(n, mosca)

  for {
    spostaMosca(n, &mosca)
    stampaStanza(n, mosca)
    time.Sleep(50 * time.Millisecond)
    cancella()
  }
}
