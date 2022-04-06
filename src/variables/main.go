package main

import "fmt"

//A linguagem Go permite declarar variáveis globais que podem ser usadas em qualquer parte do código em que são declaradas
const (
	myconstVar1 = "Minha primeira variável global constante"
	myconstVar2 = "Minha segunda variável global constante"
)

var (
	myVar1 = "Minha primeira variável global normal"
	myVar2 = "Minha segunda variável global normal"
)

func main() {
	//Go é uma linguagem tipada estaticamente, o que significa que uma variável não pode ser declarada com um tipo e depois mudá-lo
	//Entretanto, Go pode inferir o tipo da variável por meio do valor passado para ela
	//Existem três formas de declarar variáveis
	//1- Declarando nome, tipo e valor
	var msg string = "Hello World\n"
	//2- Declarando "var", nome e valor, o tipo é inferido nesse caso
	var otherstring = "O Hello World impostor\n"
	//3- Declarando nome da variável, := e o valor, o tipo é inferido aqui também
	othersstring2 := "Olha eu de novo\n"
	fmt.Printf(msg)
	fmt.Printf(otherstring)
	fmt.Printf(othersstring2)
	//Imprimindo uma mensagem que mistura variáveis
	prizemsg := "O prêmio de hoje foi %d" //Go usa interpolações para juntar diferentes tipos de variável como %s (string) e %d (int)
	prize := 9
	fmt.Printf(prizemsg, prize)
}
