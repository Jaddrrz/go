package main

import "fmt"
import "io"

func main() {
	var mm,c int
	gg := 1

	for {
		_, err := fmt.Scan(&mm)
		if err == io.EOF {
			break
		}
		if mm > 0 { 	
			c = gg
		}
		gg++
	}

	if c == 0 {
		fmt.Println("mai piovuto")
	} else {
		fmt.Println(c)
	}
}