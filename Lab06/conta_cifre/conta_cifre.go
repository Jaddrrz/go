package main

import "fmt"
import "unicode"
import "strconv"

func contaCifre(s string, numCifre *[10]int) (haCifre bool) {
	for _,v := range s {
		if unicode.IsDigit(v) {
			haCifre = true
			cifra, _ := strconv.Atoi(string(v))
			numCifre[cifra]++
		}
	}
	return
}

func main() {
	var x string
	var cifre = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var numCifre = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	c := 0

	for {
		fmt.Scan(&x)
		if x == "stop" {
			break
		}
		if contaCifre(x, &numCifre) {
			c++
		}
	}

	fmt.Println(c, "stringhe con cifre")
	fmt.Println(cifre)
	fmt.Println(numCifre)
}