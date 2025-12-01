package main

import "fmt"

func estraiPari(in []int) (out []int) {
	for _,v := range in {
		if v % 2 == 0 {
			out = append(out, v)
		}
	}
	return
}

func rimuoviMultipli(m int, in []int) (out []int) {
	for i,v := range in {
		if v % m == 0 {
			in = append(in[:i], in[i+1:]...)
		}
	}
	out = in
	return
}

func main() {
	x := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(estraiPari(x))
	fmt.Println(rimuoviMultipli(2, x))
}