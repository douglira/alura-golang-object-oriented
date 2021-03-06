package contas

import (
	"alura/2-orientacao-objetos/banco/clientes"
)

type ContaCorrente struct {
	Titular              clientes.Titular
	Agencia, NumeroConta int
	saldo                float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}

	return "saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito efetuado com sucesso", c.saldo
	}
	return "Depósito inválido", c.saldo
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) string {
	if valorDaTransferencia > 0 && valorDaTransferencia < c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return "Transferência efetuada com sucesso"
	}
	return "Transferência inválida"
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
