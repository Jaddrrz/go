package main

import "fmt" 

func main() {
	var a,b float64
	fmt.Scan(&a)
	for a != 0 {
		fmt.Scan(&b)
		if b == 0 {
			break
		}
		fmt.Println(b, "-", a, "=", b-a)
		a = b
	}
}