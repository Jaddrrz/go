package main

import "fmt"
import "math"

func main() {
	var n,closest int
	const TARGET = 8

	fmt.Println("numero:")
	fmt.Scan(&n)
	
	if n > 20 {
		fmt.Println("di nuovo:")
		fmt.Scan(&n)
	}
	closest = n
	
	for i := 1; i < 5; i++ {
		fmt.Scan(&n)
		
		if n > 20 {
			fmt.Println("di nuovo:")
			fmt.Scan(&n)
		}
		
		if math.Abs(float64(TARGET - closest)) > math.Abs(float64(TARGET - n)) {
			closest = n
		} 
	}
	fmt.Println(closest)
}