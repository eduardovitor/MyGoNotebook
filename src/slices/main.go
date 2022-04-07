package main

import (
	"fmt"
)

func main() {
	//Slices são primos dos arrays, eles podem mudar de tamanho dinamicamente, são passados por referência e permitem especificar tamanho mínimo e máximo de valores a serem guardados
	fruits := [...]string{"apples", "oranges", "bananas", "kiwis"}
	fmt.Printf("%v\n", fruits[1:3])
	fmt.Printf("%v\n", fruits[0:2])
	fmt.Printf("%v\n", fruits[:3])
	fmt.Printf("%v\n", fruits[2:])
	//Slices podem ser inicializados sem nenhum valor, mySlice := make([]int), ou:
	mySlice := make([]int, 4, 8) //O primeiro argumento indica o tipo do slice, o segundo o tamanho mínimo do slice e o terceiro o tamanho máximo
	fmt.Printf("Initial Length: %d\n", len(mySlice))
	fmt.Printf("Capacity: %d\n", cap(mySlice))
	fmt.Printf("Contents: %v\n", mySlice)
	//Slices podem armazenar valores fora da capacidade definida anteriormente, contanto que se utilize a função append()
	//Por sua flexibilidade e robustez, eles são geralmente mais usados do que arrays em Go
}
