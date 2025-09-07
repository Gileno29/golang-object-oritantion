package main

import "fmt"

type Animal struct {
	Nome  string
	Idade int
}

type Cachorro struct {
	//herança????
	/// não existe herança em golang, mas esxiste o conceito de embleding, uma astruct recebe outra dentro dela
	Animal
	Raca string
}

func (c Cachorro) EmitirSom() {
	fmt.Println("AUAU!!")
}

type Gato struct {
	Animal
	Cor string
}

func (g Gato) EmitirSom() {
	fmt.Println("MIAAAAAUU!!!!")
}

// polimorfismo

type IAnimal interface {
	EmitirSom()
}

func fazerbarulho(a IAnimal) {
	a.EmitirSom()
}

func main() {
	dog := Cachorro{
		Animal: Animal{
			Nome:  "tom",
			Idade: 12},
		Raca: "vira-lata"}
	fmt.Println(dog)

	dog.EmitirSom()
}
