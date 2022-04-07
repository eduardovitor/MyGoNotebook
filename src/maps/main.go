package main

import (
	"fmt"
	"sort"
)

func main() {
	//Maps são iguais ao conceito de dicionário, array associativo ou hash, basicamente são coleções não ordenadas de pares chave-valor
	actor := map[string]int{ //Com a função make, o map é criado. O primeiro tipo é da chave e o segundo é do valor
		"Paltrow": 43,
		"Cruise":  53,
		"Redford": 79,
		"Diaz":    43,
		"Kilmer":  56,
		"Pacino":  75,
		"Ryder":   44,
	}
	// for i := 1; i < 4; i++ { //Aqui o map é iterado
	// 	fmt.Printf("\nRUN NUMBER %d\n", i)
	// 	for key, value := range actor {
	// 		fmt.Printf("%s : %d years old\n", key, value)
	// 	}
	// }
	//Se você quiser ordenar as chaves do map, deve usar um outro tipo de dados, o slice, exemplo:
	// Armazena as chaves em um slice
	var sortedActor []string
	for key := range actor {
		sortedActor = append(sortedActor, key)
	}
	// Ordena o slice em ordem alfabética
	sort.Strings(sortedActor)
	for _, name := range sortedActor {
		fmt.Printf("%s : %d years old\n", name, actor[name])
	}
}
