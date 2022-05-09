package main

import "fmt"

type ContaCorrente struct {
	titular     string
	agencia     int
	numeroConta int
	saldo       float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito efetuado com sucesso", c.saldo
	}
	return "Depósito inválido", c.saldo
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino ContaCorrente) string {
	if valorDaTransferencia > 0 && valorDaTransferencia < c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return "Transferência efetuada com sucesso"
	}
	return "Transferência inválida"
}

func main() {
	contaDoDouglas := ContaCorrente{
		titular:     "Douglas",
		agencia:     589,
		numeroConta: 124578,
		saldo:       127.4,
	}
	contaDoGabriel := ContaCorrente{
		"Gabriel",
		589,
		784512,
		52.63,
	}
	contaDoFernando := ContaCorrente{
		titular:     "Fernando",
		agencia:     589,
		numeroConta: 542187,
	}

	fmt.Println(contaDoDouglas)
	fmt.Println(contaDoGabriel)
	fmt.Println(contaDoFernando)

	var contaDaSilvana *ContaCorrente // OU: var contaDaSilvana = new(ContaCorrente)
	contaDaSilvana = new(ContaCorrente)
	contaDaSilvana.titular = "Silvana"
	contaDaSilvana.agencia = 589

	fmt.Println(contaDaSilvana)  // &{Silvana 589 0 0}
	fmt.Println(*contaDaSilvana) // {Silvana 589 0 0}

	// ################## Comparação ########################

	contaDoAlexandre := ContaCorrente{
		titular:     "Alexandre",
		agencia:     589,
		numeroConta: 326598,
	}
	contaDoAlexandre2 := ContaCorrente{
		titular:     "Alexandre",
		agencia:     589,
		numeroConta: 326598,
	}

	/*
		TRUE: Pois o conteúdo do struct é igual
	*/
	fmt.Println(contaDoAlexandre == contaDoAlexandre2) // true
	contaDoAlexandre.saldo = 10
	/*
		FALSE: Pois um valor do struct é diferente do outro
	*/
	fmt.Println(contaDoAlexandre == contaDoAlexandre2) // false

	var contaDoPedro *ContaCorrente
	contaDoPedro = new(ContaCorrente)
	contaDoPedro.titular = "Pedro"
	contaDoPedro.agencia = 589
	contaDoPedro.numeroConta = 852963
	contaDoPedro.saldo = 250
	var contaDoPedro2 *ContaCorrente
	contaDoPedro2 = new(ContaCorrente)
	contaDoPedro2.titular = "Pedro"
	contaDoPedro2.agencia = 589
	contaDoPedro2.numeroConta = 852963
	contaDoPedro2.saldo = 250

	/*
		FALSE: Pois o ponteiro de memória de ambas são diferente
		mesmo possuindo o mesmo valor do struct
	*/
	fmt.Println(contaDoPedro == contaDoPedro2) // false
	/*
		TRUE: Desta forma é comparado o valor do struct que são iguais
	*/
	fmt.Println(*contaDoPedro == *contaDoPedro2) // true

	// ######################## Método do struct #####################

	contaDoRoberto := ContaCorrente{}
	contaDoRoberto.titular = "Roberto"
	contaDoRoberto.saldo = 500

	// var valorDoSaque float = 200
	valorDoSaque := 200.

	fmt.Println("Saldo antes do saque:", contaDoRoberto.saldo)
	fmt.Println(contaDoRoberto.Sacar(valorDoSaque))
	fmt.Println("Saldo depois do saque:", contaDoRoberto.saldo)
	status, valor := contaDoRoberto.Depositar(775)
	fmt.Println(status, valor)
	fmt.Println("Saldo depois do depósito:", contaDoRoberto.saldo)

	contaDoRoberto.Transferir(300, contaDoDouglas)
	fmt.Println("Saldo Conta Origem:", contaDoRoberto.saldo)
	fmt.Println("Saldo Conta Destino:", contaDoDouglas.saldo)
	// ################## Varidic functions ######################

	fmt.Println("Resultado da Varidic Function: ", Somando(1, 4, 32, 54, 7, 4, 2, 277))
}

func Somando(numeros ...int) int {
	resultadoDaSoma := 0
	for _, numero := range numeros {
		resultadoDaSoma += numero
	}
	return resultadoDaSoma
}
