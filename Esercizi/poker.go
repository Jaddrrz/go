package main

import "fmt"
import "strconv"

func main() {
	var n int
	var seme, valore string
	fmt.Print("Inserisci un numero da 0 a 51: ")
	fmt.Scan(&n)

	fmt.Println(n/13)
	fmt.Println("\u2666")

	switch n / 13 {
		case 0:
			seme = "\u2660"
		case 1:
			seme = "\u2663"
		case 2:
			seme = "\u2666"
		case 3:
			seme = "\u2665"
		default:
			seme = "Seme non disponibile"
	}

	switch n % 13 {
		case 0:
			valore = "A"
		case 10:
			valore = "J"
		case 11:
			valore = "Q"
		case 12:
			valore = "K"
		default:
			valore = strconv.Itoa(n % 13 + 1)
	}

	carta := valore + " " + seme

	fmt.Println("Il numero inserito corrisponde alla carta", carta)
}