package main

import "fmt"

func main() {
	string1 := `Essa é uma string que não aceita \n ou \t`
	string2 := "Essa é uma string que aceita \n e \t"
	fmt.Printf(string1)
	fmt.Printf(string2)
	//Juntando strings
	stringjoin := "O tempo"
	stringjoin2 := "piorou nos últimos dias"
	fmt.Printf(stringjoin + " " + stringjoin2)
	//Pegando o tamanho da string
	fmt.Printf("%d", len(stringjoin))
	//Imprimindo substrings
	fmt.Printf(stringjoin[0:2])
	//Comparando strings (==, >, <, <=, >=, !=)
	if stringjoin != stringjoin2 {
		fmt.Printf("String diferentes meu caro")
	}
}
