package main

import "fmt"

func main() {
	var n, prevN, c int
	fmt.Scan(&prevN)

	for {          // 1 1 1 2 2 3 3
		fmt.Scan(&n)
		if n < prevN {
			breakls
		}

		
	} 

}