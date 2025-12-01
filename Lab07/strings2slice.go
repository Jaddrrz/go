package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	input := os.Args[1:]
	
	fmt.Println(strings.Split(input[0], ","))

	fmt.Println(strings.SplitAfter(input[0], ","))

	fmt.Println(strings.Fields(input[1]))

	fmt.Println([]byte(input[2]))

	fmt.Println([]rune(input[3]))

	fmt.Println(strings.Split(input[4], ":"))
}