package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "strings"

func main() {
	var n int
	var righe []string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Quante righe? ")
	scanner.Scan()
	n,_ = strconv.Atoi(scanner.Text())

	righeScanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n && righeScanner.Scan(); i++ {
		riga := righeScanner.Text()
		righe = append(righe, riga)
	}

	patternScanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Inserire pattern: ")
		patternScanner.Scan()
		pattern := patternScanner.Text()
		c := 0
		for i := 0; i < len(righe); i++ {
			if strings.Contains(righe[i], pattern) {
				c++
			}
		}
		fmt.Println(pattern, "è contenuto in", c, "righe")
	}
}

	