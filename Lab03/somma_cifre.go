package main

import "fmt"

func main() {
	var n, original, sum int

	fmt.Scan(&n)
	original = n

	for n > 0 {
		sum += n % 10
		n /= 10
	}

	fmt.Printf("la somma delle cifre di %d è %d\n", original, sum)
}