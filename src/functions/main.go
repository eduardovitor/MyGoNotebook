package main

import (
	"fmt"
)

func OddOrEven(numero int) string {
	if numero%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
func main() {
	retorno := OddOrEven(8)
	fmt.Printf(retorno)
}
