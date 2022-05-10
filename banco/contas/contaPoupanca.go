package contas

import (
	"alura/2-orientacao-objetos/banco/clientes"
)

type ContaPoupanca struct {
	Titular                        clientes.Titular
	NumeroConta, Agencia, Operacao int
	saldo                          float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}

	return "saldo insuficiente"
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito efetuado com sucesso", c.saldo
	}
	return "Depósito inválido", c.saldo
}

func (c *ContaPoupanca) Transferir(valorDaTransferencia float64, contaDestino *ContaPoupanca) string {
	if valorDaTransferencia > 0 && valorDaTransferencia < c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return "Transferência efetuada com sucesso"
	}
	return "Transferência inválida"
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
