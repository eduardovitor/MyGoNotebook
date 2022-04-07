package main

import (
	"fmt"
)

type Person struct {
	Name string
}

type Animal struct {
	Sound string
}

//Interfaces em Go têm o mesmo conceito de Java, no entanto, a forma como tipos implementam a interface é diferente
//Para implementar uma interface, basta implementar os métodos da mesma
type Greeter interface {
	SayHi() string
}

//Tipo Pessoa implementa a interface definindo o corpo da função SayHi
func (p Person) SayHi() string {
	return "Hello! My name is " + p.Name
}

//Tipo Animal também implementa a interface definindo o corpo da função SayHi

func (a Animal) SayHi() string {
	return a.Sound
}

//Este método pode ser usado para qualquer tipo que implemente a interface Greeter

func greet(i Greeter) {
	fmt.Println(i.SayHi())
}

func main() {
	man := Person{Name: "Bob Smith"}
	dog := Animal{Sound: "Woof! Woof!"}
	fmt.Println("\nPerson : ")
	greet(man)
	fmt.Println("\nAnimal : ")
	greet(dog)
}
