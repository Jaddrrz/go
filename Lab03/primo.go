package main

import "fmt" 

func main() {
	var n int
	var isNotPrime bool
	fmt.Scan(&n)

	for i := 2; i < n; i++ {
		if n % i == 0 {
			isNotPrime = true
			fmt.Println(n, "non è primo")	
			break
		}
	}

	if !isNotPrime {
		fmt.Println(n, "è primo")
	} 
}