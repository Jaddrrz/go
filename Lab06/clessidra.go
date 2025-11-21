package main

import "fmt"
import "os"
import "os/exec"
import "strings"
import "strconv"
import "time"


func rigaClessidra(length int, sand bool, sandChar byte) string {
	if sand {
		return strings.Repeat(string(sandChar), length)
	} 
	return strings.Repeat(" ", length)
}

func clessidra(height int, time int, char byte) {
	var sand bool

	for i := 0; i < height; i++ {
		if i >= time {
			sand = true
		}
		// spazi aggiungere prima della riga, necessario per dare la forma di una clessidra
		for j := 0; j < i; j++ {       
			fmt.Print(" ")
		}
		fmt.Print("\\")
		fmt.Print(rigaClessidra(height * 2 - 1 - (i * 2), sand, char))
		fmt.Print("/\n")
	}

	sand = false
	for i := height-1; i >= 0; i-- {
		if i < time {
			sand = true
		}
		for j := i; j > 0; j-- {
			fmt.Print(" ")
		}
		fmt.Print("/")
		fmt.Print(rigaClessidra(height * 2 - 1 - (i * 2), sand, char))
		fmt.Print("\\\n")
	}
}

func cancella() {
	cmd := exec.Command("clear")
	cmd.Stdout=os.Stdout
	cmd.Run()
}

func attendi(n_seconds float64){
	time.Sleep(time.Duration(n_seconds) * time.Second)
}
 
func main() {
	var t int
	t, _ = strconv.Atoi(os.Args[1])
	cancella()

	for i := 0; i < t+1; i++ {
		clessidra(t, i, 42)
		attendi(1)
		cancella()
	}
	
}