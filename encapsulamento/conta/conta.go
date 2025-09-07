package conta

type Conta struct {
	Titular string
	Numero  string
	saldo   float64 // lower case letter makes the atribute private
}

func (c *Conta) Credito(valor float64) {
	c.saldo += valor
}

func NovaConta(titular string) *Conta {
	return &Conta{
		Titular: titular,
		Numero:  "20",
		saldo:   0.0,
	}
}
