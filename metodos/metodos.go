package main

import "fmt"

type Conta struct {
	Titular string
	Numero  string
	saldo   float64
}

// in golang methods are functions tied to structs
func (c *Conta) addCredito(valor float64) {
	c.saldo = valor

}

func newConta(titular string) *Conta {
	return &Conta{
		Titular: titular,
		Numero:  "12345",
		saldo:   0.0,
	}

}

func main() {
	conta1 := Conta{Titular: "gileno duarte",
		Numero: "123456",
		saldo:  0.0, //composição literal(criação enquanto declara)

	}

	conta2 := newConta("maria paulina")

	conta1.addCredito(20)
	fmt.Println("conta 1", conta1)
	fmt.Println("conta 2", conta2)
}
