package main

import (
	"OOexemples/clientes"
	"OOexemples/contas"
	"fmt"
)

type verificarConta interface {
	Sacar(valor float64) (string, float64)
}

func main() {
	clienteTeste := clientes.Titular{
		Nome:      "Linux",
		CPF:       "9999999",
		Profissao: "Gato",
	}
	conta1 := contas.ContaCorrente{
		Titular: clienteTeste,
	}

	clienteTeste2 := clientes.Titular{
		Nome:      "Ruby",
		CPF:       "9999999",
		Profissao: "Gato",
	}
	conta2 := contas.ContaCorrente{
		Titular: clienteTeste2,
	}

	clienteTeste3 := clientes.Titular{
		Nome:      "TESTE",
		CPF:       "1234567891011",
		Profissao: "desempregado",
	}
	conta3 := contas.ContaPoupanca{
		Titular:       clienteTeste3,
		NumeroAgencia: 11111,
		NumeroConta:   999999,
	}

	conta1.Depositar(100)
	conta2.Depositar(100)
	conta3.Depositar(1000000)

	PagaBoleto(&conta3, 91.)
	fmt.Println(conta3.ObterExtrato())

}

func PagaBoleto(conta verificarConta, valor float64) {
	conta.Sacar(valor)
}
