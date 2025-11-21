package main

import "fmt"

func main() {
	var capitale, tasso, obiettivo float64
	fmt.Println("capitale iniziale(euro): ")
	fmt.Scan(&capitale)
	fmt.Println("tasso di interesse annuale(%): ")
	fmt.Scan(&tasso)
	fmt.Println("obiettivo(euro): ")
	fmt.Scan(&obiettivo)

	anni := 0
	for capitale < obiettivo {
		capitale = capitale + capitale*(tasso/100)
		anni++
	}
	fmt.Println("Ci vogliono", anni, "per arrivare (o superare) l'obiettivo")
}