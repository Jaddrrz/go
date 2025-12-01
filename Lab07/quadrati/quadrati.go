package main

import "fmt"
import "os"
import "strconv"

func isSquare(n int) bool {
	for i := 0; i <= n; i++ {
		if i*i == n {
			return true
		}
	}
	return false
}

func main() {
	nums := os.Args[1:]
	for _, v := range nums {
		num, _ := strconv.Atoi(v)
		if isSquare(num) {
			fmt.Println(num)
		}
	}
}