package main

import (
	"fmt"
)

func main() {
	total := 0
	mean := 0
	rainfallStats := [5]int{1091, 2010, 995, 1101, 1111} //declaração de um array, também pode ser assim: var numeros[5]int
	for _, value := range rainfallStats {                //Iterando um array usando a palavra chave range
		total += value
	}
	mean = total / len(rainfallStats)
	fmt.Printf("Average rainfall: %d mm\n", mean)
}
