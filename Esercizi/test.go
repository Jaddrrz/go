package main

import "fmt"

func main() {
	var divisore, x int
	fmt.Scan(&x)
	for divisore = 2; divisore < x; divisore ++ {         
		if x % divisore == 0 {
			break
		}
	}
	if divisore == x {
		fmt.Println("è primo")
	} else {
		fmt.Println("non primo")
	}

}