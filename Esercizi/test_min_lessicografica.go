package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

type testCase struct {
	name     string
	input    string
	expected string
}

func main() {
	tests := []testCase{
		{"Caso base", "cane gatto topo", "cane 1"},
		{"Minima non prima", "gatto cane topo", "cane 2"},
		{"Più righe", "zebra\nmucca\npecora\nasino", "asino 4"},
		{"Duplicati", "pera mela mela banana", "banana 4"},
		{"Tutte uguali", "ciao ciao ciao", "ciao 1"},
		{"Maiuscole/minuscole", "albero Albero zaino", "Albero 2"},
		{"Singola parola", "ciao", "ciao 1"},
		{"Caratteri speciali", "#ciao ciao $amico", "#ciao 1"},
		{"Accenti", "mela però zucchero è", "mela 1"},
	}

	passCount := 0
	failCount := 0

	for _, tc := range tests {
		cmd := exec.Command("go", "run", "min_lessicografica.go")
		cmd.Stdin = bytes.NewBufferString(tc.input)
		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		output := string(bytes.TrimSpace(out.Bytes()))

		if err != nil {
			fmt.Printf("❌ %s: errore in esecuzione: %v\n", tc.name, err)
			failCount++
			continue
		}

		if output == tc.expected {
			fmt.Printf("✅ %s: OK\n", tc.name)
			passCount++
		} else {
			fmt.Printf("❌ %s: atteso '%s', ottenuto '%s'\n", tc.name, tc.expected, output)
			failCount++
		}
	}

	fmt.Printf("\nRisultati: %d superati, %d falliti\n", passCount, failCount)
}
