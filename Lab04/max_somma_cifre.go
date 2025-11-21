package main

import "fmt"

func main() {
	var n int

	maxSum := 0

	for {
		fmt.Scan(&n)
		if n == 999 {
			break
		}
		currSum := 0
		for n > 0 {
			currSum += n % 10
			n /= 10
		}
		if currSum > maxSum {
			maxSum = currSum
		}
	}
	fmt.Println(maxSum)

}