package main

import "fmt"

func main() {
	var n int
	fmt.Print("inserisci un numero di secondi: ")
	fmt.Scan(&n)

	gg := n / (60*60*24)
	hh := n % (60*60*24) / (60*60)
	min := n % (60*60*24) % (60*60) / 60
	sec := n % (60*60*24) % (60*60) % 60

	fmt.Print(gg, ":", hh, ":", min, ":", sec, "\n")
}