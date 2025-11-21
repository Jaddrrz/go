package main

import "fmt"
import "bufio"
import "os"

func main() {
	var text []string
	var maxLen int
	
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		line := scanner.Text()
		text = append(text, line)
	}

	for i := 0; i < len(text); i++ {
		if len(text[i]) > maxLen {
			maxLen = len(text[i])
		}
	}

	for i := 0; i < len(text); i++ {
		for j := 0; j < (maxLen - len(text[i])) / 2; j++ {
			fmt.Print(" ")
		}
		fmt.Println(text[i])
	}
}