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
}
