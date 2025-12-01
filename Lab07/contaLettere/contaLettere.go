package main

import "bufio"
import "fmt"
import "os"
import "strings"

const LEN_ALFABETO = 26

func aggiornaConteggi(s string, contaMinu *[LEN_ALFABETO]int) {
	var alfabeto [LEN_ALFABETO]string
	for i := 97; i <= 122; i++ {
		for j := 0; j < 26; j++ {
			alfabeto[j] = string(i)
		}
	}

	for _,v := range s {
		for i,lettera := range alfabeto {
			if string(v) == lettera {
				contaMinu[i]++
			}
		}
	}
}

func main() {
	var testo []string
	var contaMinu [LEN_ALFABETO]int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := scanner.Text()
		testo = append(testo, riga)
	}

	s := strings.Join(testo, " ")
	aggiornaConteggi(s, &contaMinu)
	for i := 97; i <= 122; i++ {
		for j := 0; j < 26; j++ {
			fmt.Printf("%s %d\n", string(i), contaMinu[j])
		}
	}
}