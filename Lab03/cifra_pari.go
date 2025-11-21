package main

import "fmt"

func main() {
	var n int
	var hasEven bool 
	fmt.Scan(&n)

	position := 1

	for n > 0 {
		digit := n % 10
		if digit % 2 == 0 {
			fmt.Println("la prima cifra pari", digit, "è in posizione", position)
			hasEven = true
			break
		}
		n /= 10
		position += 1
	}

	if !hasEven {
		fmt.Println("-1")
	}
}