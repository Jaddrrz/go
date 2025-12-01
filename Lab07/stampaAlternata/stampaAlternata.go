package main

import "bufio"
import "fmt"
import "io"
import "os"

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if err := scanner.Err(); err == io.EOF {
			break
		}

		riga := scanner.Text()

		input = append(input, riga)
	}
	
	for i, v := range input {
		if (i+1) % 2 == 0 {
			fmt.Println(v)
		}
	}
	for i, v := range input {
		if (i+1) % 2 != 0 {
			fmt.Println(v)
		}
	}
}