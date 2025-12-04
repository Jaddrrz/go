package main

import "bufio"
import "fmt"
import "os"

func contaVocali(s string, vocali map[rune]int) {
}

func main() {
  vocali := map[rune]int{'A'}
  //scanner := bufio.NewScanner(os.Stdin)
  //scanner.Scan()
  input := "jdhkas c'è dkasjhkdjashkdh askdh ksah @@@ €€€ ### Ħ  wi) Ø qwqwe qwyewquteuqwte q 312312 2312wweqe €łłŧŧŧŧŧ sdasdas AA JKJLKLJLKJ LIIIIII  u ù aeiou 	AEIO"
  contaVocali(input, vocali)
}
