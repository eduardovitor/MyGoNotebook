package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Pegando valores digitados na linha de comando e exibindo
	val1 := os.Args[1]
	val2 := os.Args[2]
	val3 := os.Args[3]
	val4 := os.Args[4]
	fmt.Printf("%s %s %s %s\n", val1, val2, val3, val4)
	//Checando se os valores são pares ou ímpares
	//Há funções em Go que retornam mais de um valor, nesse caso é preciso declarar duas variáveis para receber os valores
	//Se você não for usar uma variável que uma função retorna, pode colocar a empty variable (_)
	val1_converted, _ := strconv.Atoi(val1)
	if val1_converted%2 == 0 {
		fmt.Printf("%d é par", val1_converted)
	} else if val1_converted%2 != 0 {
		fmt.Printf("%d não é par", val1_converted)
	}
}
