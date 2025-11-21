package main

import "fmt"

func main() {
	var k int
	var x string
	
	fmt.Scan(&k, &x)
	k = k % 26

	for _,v := range x {
		newChar := v + rune(k)
		for newChar > 'z' {         
			newChar = 'a' + (newChar - 'z' - 1)   96   97       96 + 196 - 122 - 1      169 + 96 - 122 - 1
		}
		fmt.Print(string(newChar))
	}
	fmt.Println()
}