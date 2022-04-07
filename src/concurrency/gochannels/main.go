package main

import (
	"fmt"
	"math/rand"
	"time"
)

func broadcast(c chan int) {
	// Loop infinito para gerar números aleatórios
	for {
		/* gGera um número aleatório entre 0-999
		   e coloca no channel */
		c <- rand.Intn(999)
	}

}

func main() {
	numbersStation := make(chan int)
	// executa a função broadcast em uma thread separada
	go broadcast(numbersStation)
	// for para percorrer os valores do channel
	for num := range numbersStation {
		// Atraso para deixar a impressão dos valores mais suave
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("%d ", num)
	}

}
