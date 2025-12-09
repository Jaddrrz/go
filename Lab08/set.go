package main

import "bufio"
import "fmt"
import "os"
import "sort"
import "strings"

type StringSet map[string]struct{}

type EmptyStruct struct{}

func NewStringSet() StringSet {
	return make(StringSet)
}

func toSlice(s StringSet) []string {
	var res []string
	for k,_ := range s {
		res = append(res, k)
	}
	sort.Strings(res)
	return res
}

func Union(set1, set2 StringSet) StringSet {
	res := NewStringSet()
	for k,_ := range set1 {
		res[k] = struct{}{}
	}
	for k,_ := range set2 {
		res[k] = struct{}{}
	}
	return res
}

func Intersection(set1, set2 StringSet) StringSet {
	res := NewStringSet()
	for k,_ := range set1 {
		_, ok := set2[k] 
		if ok { 
			res[k] = struct{}{} // Se nel secondo insieme esiste una chiave uguale ad una chiave del primo, la aggiunge al risultato (res)
		}
	}
	return res
}

func Difference(set1, set2 StringSet) StringSet {
	res := NewStringSet()
	for k,_ := range set1 {
		_,ok := set2[k]
		if !ok {
			res[k] = struct{}{}
		}
	}
	return res
}

func main() {
	value := EmptyStruct{}
	var input []string
	A, B := NewStringSet(), NewStringSet()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Err() != nil {
			break
		}
		riga := strings.Fields(scanner.Text()) // "riga" corrisponde ad una slice degli elementi inseriti a stdin divisi da uno spazio
		for _,v := range riga {
			input = append(input, v)
		}
	}
	for _,v := range input {
		elem := strings.Split(v, ":") // Divido ogni elemento dell'input in base a ":" in una slice (elem), il primo elemento di queta è l'insieme, il secondo il valore da aggiungere tale insieme
		if elem[0] == "A" {
			A[elem[1]] = value
		} else {
			B[elem[1]] = value
		}
	}
	fmt.Printf("set A: %v\n", toSlice(A))
	fmt.Printf("set B: %v\n", toSlice(B))
	fmt.Printf("union: %v\n", toSlice(Union(A, B)))
	fmt.Printf("intersezione: %v\n", toSlice(Intersection(A, B)))
	fmt.Printf("differenza: %v\n", toSlice(Difference(A, B)))
}