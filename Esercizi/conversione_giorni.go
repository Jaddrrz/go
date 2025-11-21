package main

import "fmt"

func main() {
	var n int
	const GIORNI_ANNO, GIORNI_SETTIMANA = 365, 7 
	fmt.Print("Inserire un numero di giorni: ")
	fmt.Scan(&n)

	anni := n / GIORNI_ANNO
	settimane := (n % GIORNI_ANNO) / GIORNI_SETTIMANA
	giorni := (n % GIORNI_ANNO) % GIORNI_SETTIMANA
	
	fmt.Println(n, "giorni =", anni, "anni +", settimane, "settimane +", giorni, "giorni.")
}