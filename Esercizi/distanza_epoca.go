package main

import "fmt"

func main() {
	var n int
	fmt.Print("Inserisci una data (es: 7 febbraio 2013): ")
	fmt.Scan(&n)
	fmt.Println(distanzaEpoca(n))
}

func isBisestile(n int) bool {
	if n % 100 != 0 && n % 4 == 0 {
		return true
	} else if n % 100 == 0 && n % 400 == 0 {
		return true
	} else {
		return false
	}
}

func distanzaEpoca(n int) int {
	const GIORNI_NORMALE, GIORNI_BISESTILE = 365, 366
	for 
}