package main

import "fmt"

type ContaCorrente struct {
	titular     string
	agencia     int
	numeroConta int
	saldo       float64
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
}
