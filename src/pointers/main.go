package main

import "fmt"

func main() {
	a := 20
	p := &a                               // Guarda o endereço de a
	fmt.Printf("Value of a is: %d\n", *p) //Imprime o valor da variável que o ponteiro aponta
	*p = *p + 10                          //Aumenta a variável apontada pelo ponteiro em dez
	fmt.Printf("Value of a is: %d\n", *p)
}
