package main

import "fmt"

func main() {
	//Go inicializa variáveis sozinho, se você não der um valor
	integer_num := 9
	float_num := 10.7
	varbool := true
	teststring := "Cheguei chegando"
	fmt.Printf("%d %.2f %t %s", integer_num, float_num, varbool, teststring)
}
