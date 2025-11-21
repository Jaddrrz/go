package main

import "fmt"

func main() {
	var n int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	for y := n; y >= 0; y-- {                  //prima iterazione: y = 4, x = 0, 0 < 4
		for x := 0; x < y + 1; x++ {	      //                  y = 3, x = 0, 0 < 3
			fmt.Print("*")
		}
	fmt.Print("\n")
	}
}