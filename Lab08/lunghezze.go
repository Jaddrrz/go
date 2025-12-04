package main

import "bufio"
import "fmt"
import "os"
import "strings"

func aggiornaConteggio(m map[int]int, riga string) {
  words := strings.Fields(riga)
  for _, word := range words {
    m[len(word)] += 1
  }
}

func trovaChiaviMinMax(m map[int]int) (int, int) {
    var lens []int
    for i,_ := range m {
      lens = append(lens, i)
    }
    lenMin, lenMax := lens[0], lens[0]
    for i := 0; i < len(lens); i++ {
      if lens[i] < lenMin {
        lenMin = lens[i]
      } 
    }
    for i := 0; i < len(lens); i++ {
      if lens[i] > lenMax {
        lenMax = lens[i]
      } 
    }
    
    return lenMin, lenMax
}

func main() {
  count := make(map[int]int)

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    if err := scanner.Err(); err != nil {
      break
    }
    riga := scanner.Text()
    aggiornaConteggio(count, riga)
  }
  fmt.Println(count)
  for i,v := range count {
    fmt.Printf("%d,%d ", i, v)
  }
  fmt.Println()
  
  lenMin, lenMax := trovaChiaviMinMax(count)
  for i := lenMin; i <= lenMax; i++ {
    fmt.Printf("%d;%d ", i, count[i])
  }
  fmt.Println()
  for i := lenMin; i <= lenMax; i++ {
    if count[i] == 0 {
      continue
    }
    fmt.Printf("%d;%d ", i, count[i])
  }
  fmt.Println()
}
