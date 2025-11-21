package main

import "fmt"

func main() {
	var n, m, sum int
	fmt.Scan(&n)
	if n != -1 {
		sum = n
	}

	for n != -1 {
		m = n
		fmt.Scan(&n)
		
		if n == -1 {
			break
		}
		
		if n < m {
			fmt.Println("-")
		} else {
			fmt.Println("+")
		}
		sum += n
	}
	fmt.Println("somma:", sum)
}