package contas

import (
	"OOexemples/clientes"
	"OOexemples/extrato"
	"OOexemples/date"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
	Extrato       extrato.Extrato
}

func (c *ContaCorrente) Sacar(valor float64) (string, float64) {
	podeSacar := valor > 0 && valor <= c.saldo
	if podeSacar {
		c.saldo -= valor
		c.Extrato.Eventos = append(c.Extrato.Eventos, "Saque realizado com sucesso.", date.DateTime())
		return "Saque realizado com sucesso.", c.saldo
	} else {
		c.Extrato.Eventos = append(c.Extrato.Eventos, "Tentativa de saque.", date.DateTime())
		return "saldo insuficiente.", c.saldo
	}

}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	podeDepositar := valor > 0
	if podeDepositar {
		c.saldo += valor
		c.Extrato.Eventos = append(c.Extrato.Eventos, "Deposito realizado com sucesso.", date.DateTime())
		return "Deposito realizado com sucesso.", c.saldo
	} else {
		c.Extrato.Eventos = append(c.Extrato.Eventos, "Tentativa de deposito realizado.", date.DateTime())
		return "Valor precisa ser maior que zero.", c.saldo
	}

}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	if c.saldo > 0 && valor <= c.saldo && valor > 0 {
		c.saldo -= valor
		contaDestino.Depositar(valor)
		c.Extrato.Eventos = append(c.Extrato.Eventos, "Transferencia realizada com sucesso.", date.DateTime())
		return true
	} else {
		c.Extrato.Eventos = append(c.Extrato.Eventos, "Houve uma tetativa de transferencia.", date.DateTime())
		return false
	}

}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}

func (c *ContaCorrente) ObterExtrato() []string {
	return c.Extrato.Eventos
}

// func date.DateTime() string {
// 	dt := time.Now()
// 	return dt.String()
// }

/*
aqui é como criamos um método de um onjeto, a sintaxe (c. <struct>) nos mostra exatamente isso.
a palavrinha c. é como se fosse self.

para determinar se um valor será "publico" ou "privado" usamos Maiusculo (public) e minusculo (private)
*/
