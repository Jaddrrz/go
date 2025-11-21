package main

import "fmt"
import "unicode"

func main() {
	var x string
	var sameCase bool
	fmt.Scan(&x)
	
	for i:=0; i<len(x)-1; i++ {
		if unicode.IsLower(rune(x[i])) == unicode.IsLower(rune(x[i+1])) {
			sameCase = true
		} else {
			fmt.Println("sia minuscole che maiuscole")
			sameCase = false
			break
		}
	}
	if sameCase && unicode.IsLower(rune(x[0])) {
		fmt.Println("solo minuscole")
	} else if sameCase && unicode.IsUpper(rune(x[0])) {
		fmt.Println("solo maiuscole")
	}
}