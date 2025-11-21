package main

import "fmt"
import "math"

func main() {
	var somma, n int
	var radice float64
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	for n > 0 {
		radice = math.Sqrt(float64(n)) 
		fmt.Println(n, radice, int(math.Pow(float64(int(radice)), 2.0)))
		if n == int(math.Pow(float64(int(radice)), 2.0)) {
			somma += n
			fmt.Println(n, somma)
		}
		n--
	}
	fmt.Println(somma)
}