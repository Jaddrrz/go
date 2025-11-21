package main

import "fmt"

func main() {
	var x string

	fmt.Scan(&x)

	for _, r := range x {
		if string(r) >= "a" && string(r) <= "z" {
			fmt.Println(string(r), "\t", "minuscolo")
		} else {
			fmt.Println(string(r), "\t", "maiuscolo")
		}
	}
}