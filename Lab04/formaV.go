package main

import "fmt"

func main() {
	var n int	
	fmt.Scan(&n)

	for y:=1; y<=n; y++ {
		for x:=1; x<=n*2; x++ {
			if (y==x && x<=n) || (y+x==2*n && x>n) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}