package main

import (
	"alura/2-orientacao-objetos/banco/contas"
	utilitarios "alura/2-orientacao-objetos/banco/utils"
	"fmt"
)

func main() {
	contaDoDouglas := contas.ContaCorrente{
		Titular:     "Douglas",
		Agencia:     589,
		NumeroConta: 124578,
		Saldo:       127.4,
	}
	contaDoGabriel := contas.ContaCorrente{
		"Gabriel",
		589,
		784512,
		52.63,
	}
	contaDoFernando := contas.ContaCorrente{
		Titular:     "Fernando",
		Agencia:     589,
		NumeroConta: 542187,
	}

	fmt.Println(contaDoDouglas)
	fmt.Println(contaDoGabriel)
	fmt.Println(contaDoFernando)

	var contaDaSilvana *contas.ContaCorrente // OU: var contaDaSilvana = new(contas.ContaCorrente)
	contaDaSilvana = new(contas.ContaCorrente)
	contaDaSilvana.Titular = "Silvana"
	contaDaSilvana.Agencia = 589

	fmt.Println(contaDaSilvana)  // &{Silvana 589 0 0}
	fmt.Println(*contaDaSilvana) // {Silvana 589 0 0}

	// ################## Comparação ########################

	contaDoAlexandre := contas.ContaCorrente{
		Titular:     "Alexandre",
		Agencia:     589,
		NumeroConta: 326598,
	}
	contaDoAlexandre2 := contas.ContaCorrente{
		Titular:     "Alexandre",
		Agencia:     589,
		NumeroConta: 326598,
	}

	/*
		TRUE: Pois o conteúdo do struct é igual
	*/
	fmt.Println(contaDoAlexandre == contaDoAlexandre2) // true
	contaDoAlexandre.Saldo = 10
	/*
		FALSE: Pois um valor do struct é diferente do outro
	*/
	fmt.Println(contaDoAlexandre == contaDoAlexandre2) // false

	var contaDoPedro *contas.ContaCorrente
	contaDoPedro = new(contas.ContaCorrente)
	contaDoPedro.Titular = "Pedro"
	contaDoPedro.Agencia = 589
	contaDoPedro.NumeroConta = 852963
	contaDoPedro.Saldo = 250
	var contaDoPedro2 *contas.ContaCorrente
	contaDoPedro2 = new(contas.ContaCorrente)
	contaDoPedro2.Titular = "Pedro"
	contaDoPedro2.Agencia = 589
	contaDoPedro2.NumeroConta = 852963
	contaDoPedro2.Saldo = 250

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

	contaDoRoberto := contas.ContaCorrente{}
	contaDoRoberto.Titular = "Roberto"
	contaDoRoberto.Saldo = 500

	// var valorDoSaque float = 200
	valorDoSaque := 200.

	fmt.Println("Saldo antes do saque:", contaDoRoberto.Saldo)
	fmt.Println(contaDoRoberto.Sacar(valorDoSaque))
	fmt.Println("Saldo depois do saque:", contaDoRoberto.Saldo)
	status, valor := contaDoRoberto.Depositar(775)
	fmt.Println(status, valor)
	fmt.Println("Saldo depois do depósito:", contaDoRoberto.Saldo)

	fmt.Println("Antes da transferencia:", contaDoRoberto.Saldo, contaDoDouglas.Saldo)
	contaDoRoberto.Transferir(300, &contaDoDouglas)
	fmt.Println("Saldo Conta Origem:", contaDoRoberto.Saldo)
	fmt.Println("Saldo Conta Destino:", contaDoDouglas.Saldo)
	// ################## Varidic functions ######################

	fmt.Println("Resultado da Varidic Function: ", utilitarios.Somando(1, 4, 32, 54, 7, 4, 2, 277))
}
