package main

import "fmt"

func isTwinPrime(x, y int) bool {
	return isPrime(x) && isPrime(y) && isDifferenceTwo(x, y)
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

func isDifferenceTwo(x, y int) bool {
	var n1, n2 int
	if x > y {
		n1, n2 = x, y
	} else {
		n1, n2 = y, x
	}
	fmt.Println(n1 - n2)
	return (n1 - n2 == 2)
}

func main() {
	var x, y int
	fmt.Print("Inserisci due numeri interi: ")
	fmt.Scan(&x, &y)
	if isTwinPrime(x,y) {
		fmt.Println("I due numeri sono primi gemelli")
	}
}