package main

import "fmt"

func main() {
	var n, c int

	maxN := 0

	for i := 0; i < 10; i++ {
		fmt.Scan(&n)
		currN := n 
		if currN > maxN {
			maxN = currN
			c = 1
		} else if currN == maxN {
			c++
		}
	}
	
	fmt.Println(maxN, c)
}