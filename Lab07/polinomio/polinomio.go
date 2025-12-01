package main

import "bufio"
import "fmt"
import "math"
import "os"
import "strconv"
import "strings"

func main() {
	var sum, x float64
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	valori := strings.Fields(input)
	
	x, _ = strconv.ParseFloat(valori[len(valori)-1], 64)
	for i := 0; i < len(valori)-1; i++ {
		coefficiente, _ := strconv.ParseFloat(valori[i], 64)
		//exp, _ := strconv.ParseFloat(i, 64)
		sum += coefficiente * math.Pow(x, float64(i))
	}

	fmt.Println(sum)
}