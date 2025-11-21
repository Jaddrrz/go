package main

import "fmt"

func main() {
	var n string
	
	c := 0

	for {
		fmt.Scan(&n)
		newC := 0
		if n == "000" {
			break
		}
		for _,v := range n {
			if (v-'0') % 2 == 0 {
				newC++   //numero cifre pari
			} 
		}
		if newC > c {
			c = newC
		}
	}

	fmt.Println(c)
}