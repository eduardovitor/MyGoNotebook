package main

import (
	"fmt"
)
//Encapsulamento em Go se refere a visibilidade ou não de um pacote, como também à visibilidade de funções/variáveis dentro daquele pacote
//Para tornar acessar um pacote, deve-se importá-lo
/*
Para deixar uma variável pública, coloca-se a primeira letra da definição dela como maiúscula, exemplo:
type Account struct {
    UserID      int    // exported
    accessToken string // unexported
}
*/

type rect struct {
	height int
	width  int
}

type Discount struct {
	percent float32

	promotionId string
}

type ManagersSpecial struct {
	Discount //Tipos podem ser colocados dentro de outro tipo, nesse caso, não é preciso dizer nome de variável, para inicializar essa struct use:  managerSpecial := ManagersSpecial{januarySale, 10.00}, sendo januarySale do tipo Discount
	extraoff float32
}

//Structs podem ser usadas da mesma forma que classes de outras linguagens como Java ou Python
//Structs são passadas por cópia(valor), para modificar os valores de forma permanente devem ser utilizados ponteiros
func main() {
	r := rect{height: 12, width: 20}
	fmt.Printf("Height: %d\n", r.height)
	fmt.Printf("Width: %d\n", r.width)
	fmt.Printf("Area: %d\n", r.area())
	fmt.Printf("\nDouble it!\n\n")
	r.double()
	fmt.Printf("Height: %d\n", r.height)
	fmt.Printf("Width: %d\n", r.width)
	fmt.Printf("Area: %d\n", r.area())
}

func (r rect) area() int {
	h := r.height
	w := r.width
	return h * w
}

func (r *rect) double() {
	r.height *= 2
	r.width *= 2
}
