package main

import "fmt"
import "io"
import "strings"
import "strconv"

func tot(p, q, s float64) float64 {
	return p*q - p*q*s
}

func main() {
	var prezzo, quantità, sconto, prezzoFinale float64

	for {
		var input string

		_, err := fmt.Scan(&input)
		if err == io.EOF {
			break
		}

		valori := strings.Split(input, "#")
		prezzo, _ = strconv.ParseFloat(valori[0], 64) 
		quantità, _ = strconv.ParseFloat(valori[1], 64)
		sconto, _ = strconv.ParseFloat(valori[2], 64)

		prezzoFinale += tot(prezzo, quantità, sconto)
	}
	fmt.Println(prezzoFinale)
}