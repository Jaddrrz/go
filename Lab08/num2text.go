package main

import "fmt"

func main() {
  var n, word string
 
  nums := make(map[string]bool)
  nums2text := make(map[string]string)
  
  fmt.Print("un numero: ")
  fmt.Scan(&n)
  for _,v := range n {
    nums[string(v)] = true
  }

  for num,_ := range nums {
    fmt.Printf("parola per %s ? ", num)
    fmt.Scan(&word)
    nums2text[num] = word
  }

  for i := 0; i < len(n)-1; i++ {  
    fmt.Printf("%s - ", nums2text[string(n[i])])
  }
  fmt.Printf("%s\n", nums2text[string(n[len(n)-1])])
}  
