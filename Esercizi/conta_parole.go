package main 

import "fmt"

func main() {
	var x string
	x = "Abbiamo visto un bellissimo film ieri sera"
	counter := 0

	for _,v := range x {
		if string(v) == " " {
			counter += 1
		}
	}
	counter += 1
	fmt.Println("parole:", counter)
}