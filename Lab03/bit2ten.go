package main

import "fmt"
import "math"

func main() {
	var n, sum int
	var exp float64
	fmt.Scan(&n)
	
	for n > 0 {
		if n % 10 == 1 {
			sum += int(math.Pow(2.0, exp))
		}
		n /= 10
		exp += 1
	}
	fmt.Println(sum)
}