package main

import "bufio"
import "fmt"
import "os"
import "strings"
import "slices"

type Capoluogo struct {
	Nome, Sigla, Regione, Superficie string
}

func main() {
  	capoluoghi := make(map[string]Capoluogo)

	input := os.Args[1:]
	
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var capoluogo Capoluogo
		riga := strings.Split(scanner.Text(), ",")
		riga = slices.Delete(slices.Delete(riga, 3, 4), 4,6) // rimuove valori da non utilizzare
		capoluogo = Capoluogo{riga[0], riga[1], riga[2], riga[3]}
		capoluoghi[riga[1]] = capoluogo
	}
	for _,v := range input {
		fmt.Println(capoluoghi[v])
	}	
}