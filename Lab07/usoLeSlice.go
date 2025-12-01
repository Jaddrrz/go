/*
Scrivere un programma usoLeSlice.go che

1. legge da linea di comando una lista di almeno 3 parole, le salva in una slice e stampa la slice;
2. legge da standard input un'altra lista di parole (anche almeno 3), le salva in una slice e stampa la slice.

Quindi fa le seguenti operazioni e man mano stampa il risultato ottenuto:

 3. crea un nuova lista in cui salva gli elementi della prima lista seguiti dagli elementi della seconda lista
 4. mette in ordine alfabetico (lessicografico) la slice (usare una funzione di libreria) e stampa la slice ottenuta
 5. cancella l'ultima parola dalla slice
 6. rimuove dalla slice gli elementi di indice dal 2 (incluso) al 4 (escluso)
 7. crea una nuova slice con "aa", "bb" e "cc"
 8. inserisce la nuova slice in posizione 1 della vecchia slice ("aa" avrà indice 1, ecc. e gli elementi della vecchia slice, tranne quello di indice 0, andranno spostati a destra della slice inserita)
 9. legge una nuova parola e la inserisce in posizione 2 della slice
10. legge una nuova parola e la inserisce in fondo alla slice
11. estende la slice con una nuova slice di lunghezza 2 (di stringhe vuote)
12. inserisce alla posizione 3 una nuova slice di lunghezza 3 (di stringhe vuote)
13. copia in una nuova slice la slice ottenuta fin qui
14. dalla copia rimuove l'ultimo elemento e stampa sia la vecchia slice che la copia.

Completare il programma inserendo al posto dei "..." le istruzioni occorrenti.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"slices"
	// ...
)

func main() {
	//1. legge da linea di comando una lista di almeno 3 parole e le salva in una slice
	if len(os.Args) < 4 {
		fmt.Println("almeno 3 parole su CL")
		return
	}
	// ...
  	wordsFromCL := os.Args[1:]
	fmt.Println(" 1.", wordsFromCL)

	///2. legge da standard input un'altra lista di parole (anche almeno 3) e le salva in una slice
	fmt.Println("scrivi almeno 3 parole (seguite da ctrl D)")
	// ...
	var wordsFromStdin []string
	scanner1 := bufio.NewScanner(os.Stdin)
	for scanner1.Scan() {
		riga := scanner1.Text()
		wordsFromStdin = append(wordsFromStdin, riga)
	}	
	fmt.Println(" 2.", wordsFromStdin)

	//3. crea un nuova lista mySlice in cui salva gli elementi della prima lista seguiti dagli elementi della seconda lista
	// ...
	mySlice := make([]string, len(wordsFromCL))
	copy(mySlice, wordsFromCL)
	for _,v := range wordsFromStdin {
		mySlice = append(mySlice, v)
	}

	fmt.Println(" 3.", mySlice)

	// 4. mette in ordine alfabetico (lessicografico) la slice (usare una funzione di libreria) e stampa la slice ottenuta
	// ...
	sort.Strings(mySlice)
	fmt.Println(" 4.", mySlice)

	// 5. cancella l'ultima parola dalla slice
	// ...

	mySlice = slices.Delete(mySlice, len(mySlice)-1, len(mySlice))
	fmt.Println(" 5.", mySlice)

	// 6. rimuove dalla slice gli elementi di indice dal 2 (incluso) al 4 (escluso)
	// ...
	i := 2
	j := 4
	mySlice = slices.Delete(mySlice, i, j)
	fmt.Println(" 6.", mySlice)

	// 7. crea una nuova slice con "aa", "bb" e "cc" e la stampa
	// ...
	myNewSlice := []string{"aa", "bb", "cc"}
	fmt.Println(" 7.", myNewSlice)

	// 8. inserisce la nuova slice in posizione 1 della vecchia slice
	// ...
	mySlice = append(mySlice[:1], append(myNewSlice, mySlice[1:]...)...)
	fmt.Println(" 8.", mySlice)

	// 9. legge una nuova parola e la inserisce in posizione 2 della slice
	fmt.Print("scrivi una parola: ")
	// ...
	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Scan()
	w1 := scanner2.Text()
	mySlice = slices.Insert(mySlice, 2, w1)
	fmt.Println(" 9.", mySlice)

	//10. legge una nuova parola e la inserisce in fondo alla slice
	fmt.Print("scrivi una parola: ")
	// ...
	scanner2.Scan()
	w2 := scanner2.Text()
	mySlice = append(mySlice, w2)
	fmt.Println("10.", mySlice)

	// 11. estende la slice con una nuova slice di lunghezza 2 (di stringhe vuote)
	// ...
	mySlice = append(mySlice, make([]string, 2)...)
	fmt.Println("11.", mySlice)

	// 12. inserisce alla posizione 3 una nuova slice di lunghezza 3 (di stringhe vuote)
	// ...
	mySlice = append(mySlice[:3], append(make([]string, 2), mySlice[3:]...)...)
	fmt.Println("12.", mySlice)

	// 13. copia in una nuova slice la slice ottenuta fin qui
	// ...
	copyOfMySlice := make([]string, len(mySlice))
	copy(copyOfMySlice, mySlice)
	fmt.Println("13.", copyOfMySlice)

	// 14. dalla copia rimuove l'ultimo elemento e stampa sia la vecchia slice che la copia.
	// ...
	copyOfMySlice = slices.Delete(copyOfMySlice, len(copyOfMySlice)-1, len(copyOfMySlice))
	fmt.Println("14.", mySlice, "\n   ", copyOfMySlice)
}

