package conta

import (
	"fmt"
	"object-orientation/encapsulamento/conta"
)

func main() {
	conta1 := conta.Conta{Titular: "gileno duarte", Numero: "888888", saldo: 10.0}
	conta2 := conta.NovaConta("maria")

	fmt.Println("conta 1", conta1)
	fmt.Println("Conta 2", conta2)

	conta2.Credito(float64(10.0))
	fmt.Println("saldo conta 2", conta2)

}
