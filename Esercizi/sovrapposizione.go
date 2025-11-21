package main

import "fmt"

func main() {
	var gg1, gg2, start1, start2, end1, end2 int
	
	fmt.Print("Primo appuntamento: ")
	fmt.Scan(&gg1, &start1, &end1)

	fmt.Println("Secondo appuntamento: ")									// ------               --------     ------------            ----
	fmt.Scan(&gg2, &start2, &end2)                                          //     ------	     --------             -----          --------------

	if (start1 < start2 && end1 > start2) || (start1 > start2 && end1 > end2) || (start1 < start2 && end1 > end2) || (start1 > start2 && end1 < end2) {
		fmt.Println("si sovrappongono")	
	} else {
		fmt.Println("non si sovrappongono")
	}
	/*var giorno1, giorno2 int
	var inizio1, fine1, inizio2, fine2 float64

	fmt.Print("Inserisci giorno, ora inizio e ora fine del primo appuntamento: ")
	fmt.Scan(&giorno1, &inizio1, &fine1)

	fmt.Print("Inserisci giorno, ora inizio e ora fine del secondo appuntamento: ")
	fmt.Scan(&giorno2, &inizio2, &fine2)

	if giorno1 == giorno2 && !(fine1 <= inizio2 || fine2 <= inizio1) {
		fmt.Println("Gli appuntamenti si sovrappongono.")
	} else {
		fmt.Println("Gli appuntamenti non si sovrappongono.")
	}*/
	
}