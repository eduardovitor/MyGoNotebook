package main

import "fmt"

func main() {
	//Existe apenas o for como estrutura de repetição em Go. Porém existem formas diferentes de usá-lo que substituem a necessidade de um while ou do-while
	//Primeira forma
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("The number is even")
		} else {
			fmt.Printf("The number is odd")
		}
	}
	//Segunda forma
	i2 := 0
	for i2 < 10 {
		fmt.Printf("%d", i2)
		i2 = i2 + 1
	}
	//Terceira forma
	i3 := 0
	for {
		if i3 == 5 {
			break
		}
		fmt.Printf("%d", i3)
		i3++
	}
}
