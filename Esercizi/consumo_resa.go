package main

import "fmt"

func main() {
	var distanza, carburante float64
	fmt.Print("Inserisci la distanza percorsa (in km) e la quantià di carburante (in l) di un motore: ")
	fmt.Scan(&distanza, &carburante)

	consumo_medio := carburante / distanza
	resa := distanza / carburante

	fmt.Println("Consumo medio:", consumo_medio, "l/km")
	fmt.Println("Resa media:", resa, "km/l")
}