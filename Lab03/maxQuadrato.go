package main

import "fmt" 

func main() {
	var n, square int
	fmt.Scan(&n)


	for i := 0; i*i < n ; i++ {
		square = i*i		
	}

	fmt.Println(square)
}