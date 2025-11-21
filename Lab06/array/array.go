package main

import "fmt"

const DIM = 5

func reverse(a *[5]int) {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a) - i - 1] = a[len(a) - i - 1], a[i]
	}
}

func switchFirstLast(a *[5]int) {
	a[0], a[len(a) - 1] = a[len(a) - 1], a[0]
}

func main() {
	var numeri = [5]int{0,1,2,3,4}

	fmt.Println(numeri)
	reverse(&numeri)
	fmt.Println(numeri)
	switchFirstLast(&numeri)
	fmt.Println(numeri)
}