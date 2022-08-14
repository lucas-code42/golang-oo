package contas

import (
	"OOexemples/clientes"
	"OOexemples/date"
	"OOexemples/extrato"
)

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
	Extrato       extrato.Extrato
}

func (c *ContaPoupanca) Sacar(valor float64) (string, float64) {
	podeSacar := valor > 0 && valor <= c.saldo
	if podeSacar {
		c.saldo -= valor
		c.Extrato.Eventos = append(c.Extrato.Eventos, "Saque realizado com sucesso", date.DateTime())
		return "Saque realizado com sucesso", c.saldo
	} else {
		c.Extrato.Eventos = append(c.Extrato.Eventos, "Tentativa de saque.", date.DateTime())
		return "Saldo insuficiente", c.saldo
	}
}

func (c *ContaPoupanca) Depositar(valor float64) (string, float64) {
	podeDepositar := valor > 0
	if podeDepositar {
		c.saldo += valor
		c.Extrato.Eventos = append(c.Extrato.Eventos, "Deposito realizado com sucesso", date.DateTime())
		return "Depósito realizado com sucesso!", c.saldo
	} else {
		c.Extrato.Eventos = append(c.Extrato.Eventos, "Tentativa de depósito", date.DateTime())
		return "Valor precisa ser maior que zero", c.saldo
	}
}

func (c *ContaPoupanca) ObterExtrato() []string {
	return c.Extrato.Eventos
}
