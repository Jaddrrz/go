package main

import "fmt"

func isPrime13(x int) bool {
	return isPrime(x) && isSum13(x) 
}

func isPrime(x int) bool {
	var divisore int
	for divisore = 2; divisore < x; divisore ++ {         
		if x % divisore == 0 {
			break
		}
	}
	return divisore == x
}

func isSum13(x int) bool {
	var somma int
	for x != 0 {
		somma += x % 10
		x /= 10
	}
	return somma == 13
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := n; i > 0; i-- {
		if isPrime13(i) {
			fmt.Println(i)
		}
	}
}