package main

import "fmt"
import "math/rand"

func main() {
	var n, x int
	var indovinato bool
	const MAX = 10

	x = rand.Intn(MAX)
	attempts := 1

	fmt.Print("indovina il numero: ")
	fmt.Scan(&n)
	
	for attempts <= (MAX / 2) && indovinato == false {
		if n > MAX || n == 0 {
			fmt.Print("fuori intervallo! ")
			fmt.Scan(&n)
		} else if n == x {
			fmt.Println("indovinato in", attempts, "tentativi")
			indovinato = true
		} else {
			fmt.Print("riprova: ")
			fmt.Scan(&n)
			attempts++
		}
	}
	
	if indovinato == false {
		fmt.Println("hai perso, il numero era", x)
	}
}