package main

import "fmt"
import "io"

func main() {
	var curr, min string
	var minLen, posMin int

	c := 1
	
	for {
		_, err := fmt.Scan(&curr)
		if err == io.EOF {
			break
		} 

		if c == 1 {                // per la prima iterazione 
			minLen = len(curr)
			min = curr
			posMin = c
			c++
			continue
		} else if len(curr) < len(min) {   // trovare la lunghezza minima tra parola corrente e minima da utilizzare in seguito per l'elaborazione
			minLen = len(curr)
		} else {
			minLen = len(min)
		}
	
		for i := 0; i < minLen; i++ {  // compara lettera per lettera le 2 parole
			if curr[i] < min[i] {
				min = curr
				posMin = c
			} else {
				break
			}
		}

		c++
	}
	fmt.Println(min, posMin)
}