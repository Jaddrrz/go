package main

import "bufio"
import "fmt"
import "os"
import "strings"
import "slices"

func main() {
  	regioni := make(map[string][]string)

	input := os.Args[1:]
	
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := strings.Split(scanner.Text(), ",")
		riga = slices.Delete(slices.Delete(riga, 1, 2), 2, 5) // rimuove valori da non utilizzare
		regioni[riga[1]] = append(regioni[riga[1]], riga[0]) // aggiunge alla slice corrispondente alla regione, la provincia
	}
	
	for _,v := range input {
		fmt.Printf("%s: %v\n", v, regioni[v])
	}	
}