package main

import (
   "fmt"
   "strings"
   "strconv"
)

func main() {
   var dataGMA string
   fmt.Print("data nel formato giorno/mese/anno: ")
   fmt.Scan(&dataGMA)
   g, m, a := stringGMA2GMA(dataGMA)
   fmt.Println("giorno", g)
   fmt.Println("mese", num2mese(m))
   fmt.Println("anno", a)

}

func stringGMA2GMA(data string) (int, int, int) {
	primaBarra := strings.Index(data, "/")
	secondaBarra := primaBarra + strings.Index(data[primaBarra + 1:], "/") + 1

	g, _ := strconv.Atoi(data[:primaBarra])                      
	m, _ := strconv.Atoi(data[primaBarra + 1: secondaBarra])
	a, _ := strconv.Atoi(data[secondaBarra + 1:])
   return g, m, a
}

func num2mese(m int) string {
   var mesi = [13]string{"", "gen", "feb", "mar", "apr", "mag", "giu", "lug", "ago", "set", "ott", "nov", "dic"}
   return mesi[m]
}
