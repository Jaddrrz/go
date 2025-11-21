package main

import "bufio"
import "fmt"
import "strconv"
import "strings"
import "time"
import "os"

type Orario struct {
	h, min, sec int
}

func estraiOrario(s string) (orario Orario) {
	primoSpazio := strings.Index(s, " ")
	secondoSpazio := primoSpazio + strings.Index(s[primoSpazio+1:], " ") + 1
	
	orario.h,_ = strconv.Atoi(s[:primoSpazio])
	orario.min,_ = strconv.Atoi(s[primoSpazio+1:secondoSpazio])
	orario.sec,_ = strconv.Atoi(s[secondoSpazio+1:])
	
	return 
}

func countdown(o *Orario) {
	if o.sec == 0 {
		o.sec = 59
		downMin(o)
	} else {
		o.sec--
	}
}

func downMin(o *Orario) {
	if o.min == 0 {
		o.min = 59
		downHour(o)
	} else {
		o.min--
	}
}

func downHour(o *Orario) {
	o.h--
}

func main() {
	var x string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("time(hh mm ss): ")
	scanner.Scan()
	x = scanner.Text()

	orario := estraiOrario(x)
	
	for {
		countdown(&orario)
		fmt.Println(orario)
		time.Sleep(time.Duration(1) * time.Second)
	}
}