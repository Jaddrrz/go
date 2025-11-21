package main

import "fmt"
import "math"

func main() {
	var x1, y1, m, q float64
	const epsilon = 1e-6

	fmt.Scan(&x1, &y1, &m, &q)

	y := m*x1 + q

	if math.Abs(y1 - y) < epsilon {
		fmt.Println("sulla retta")
	} else if y1 > y {
		fmt.Println("sopra")
	} else {
		fmt.Println("sotto")
	}
}