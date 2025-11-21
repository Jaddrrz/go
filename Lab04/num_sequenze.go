package main

import "fmt"

func main() {
	var n, c int

	for {
		prevN := n
		fmt.Scan(&n)
		
		if n == 2 {
			break
		} else if n == 0 && prevN == 1 {
			c++
		}
	}
	fmt.Println(c)
}