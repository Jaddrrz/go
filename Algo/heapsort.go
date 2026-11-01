package main

import "fmt"
import "github.com/ema-ts01/rand_list"
import "time"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func rearrange(a []int, r int, l int) {
	v := r
	x := a[v]
	toCollocate := true
	for {
		if 2 * v + 1 >= l {
			toCollocate = false
		} else {
			u := 2 * v + 1
			if u + 1 < l && a[u + 1] > a[u] {
				u += 1
			}
			if a[u] > x {
				a[v] = a[u]
				v = u
			} else {
				toCollocate = false
			}
		}
		if !toCollocate {
			break
		}
	}
	a[v] = x
}

func createHeap(a []int) {
	for i := len(a) - 1; i >= 0; i-- {
		rearrange(a, i, len(a))
	}
}

func heapSort(a []int) {
	createHeap(a)
	for l := len(a) - 1; l >= 1; l-- {
		a[0], a[l] = a[l], a[0]
		rearrange(a, 0, l)
	}
}

func main() {

	var dim, max int
	fmt.Println("Dimensione lista:")
	fmt.Scan(&dim)
	fmt.Println("Numero massimo della lista:")
	fmt.Scan(&max)
	list := rand_list.Generate(max, dim)
	fmt.Println("lista originale:\t", list)
	start := time.Now()
	heapSort(list)
	fmt.Println("lista ordinata:\t\t", list)
	defer func() {
		fmt.Printf("Esecuzione main terminata in: %v\n", time.Since(start))
	}()
}

