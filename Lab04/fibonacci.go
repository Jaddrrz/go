package main

import "fmt"

func main() {
	var n, prev1, prev2, curr int
	fmt.Scan(&n)
	if n >= 2 {
		prev1, prev2, curr = 1, 1, 2
		fmt.Println("*")
		fmt.Println("*")
		for i := 2; i < n; i++ {
			curr = prev1 + prev2
			prev1 = prev2
			prev2 = curr
			for j := 0; j < curr; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	} else if n <= 0 {
		fmt.Println("Numero invalido")
	} else if n == 1 {
		fmt.Println("*")
	}
}