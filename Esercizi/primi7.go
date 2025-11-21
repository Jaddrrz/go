package main

import "fmt"

func isPrime7(x int) bool {
	return isPrime(x) && isLastDigit7(x)
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

func isLastDigit7(x int) bool {
	return x % 10 == 7
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := n; i >= 7; i-- {
		if isPrime7(i) {
			fmt.Println(i)
		}
	}
}