package main

import "fmt"
import "math/rand"

//const SEMI, VALORI, MAZZO = 4, 13, 52

type Carta struct {
	valore, seme string
}

func carta(n int) (x Carta, ok bool) {
	if n < 0 || n > 52 {
		ok = false
	} 

	valori := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	semi := []string{"\u2665", "\u2666", "\u2663", "\u2660"}

	x.valore = valori[n%13]
	x.seme = semi[n/13]
	ok = true

	return 
}

func estraiCarta() Carta {
	n := rand.Intn(52)
	carta, _ := carta(n)
	return carta
}

func dai4Carte() [4]Carta {
	var carte [4]Carta
	for i := 0; i < 4; i++ {
		carte[i] = estraiCarta()
	}
	return carte
}

/*func main() {
	fmt.Println(carta(51))
	x := estraiCarta()
	fmt.Println(x)
	mano := dai4Carte()
	fmt.Println(mano)
}*/