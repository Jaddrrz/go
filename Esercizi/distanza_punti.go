package main

import "fmt"
import "math"

func main() {
	var x1, y1, x2, y2 float64
	fmt.Print("Inserisci le coordinate di un punto: ")
	fmt.Scan(&x1, &y1)

	fmt.Println("Inserisci le coordinate di un secondo punto: ")
	fmt.Scan(&x2, &y2)

	distanza := math.Sqrt((x2 - x1)*(x2 - x1) + (y2 - y1)*(y2 - y1))
	fmt.Println("La distanza tra i 2 punti è:", distanza)
}