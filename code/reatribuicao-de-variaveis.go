package main

import (
	"fmt"
	"os"
)

func main() {
	/* Reatribuição de variávels com := */

	var d int = 1
	d := 1

	// equivale a:

	var d int = 1
	var d int = 1

	// Ambas operações são ilegais

	fmt.Println(d)

	// Reatribuindo corretamente um valor a variável err */
	f, err := os.Open("/tmp/reasign")
	s, err := f.Stat()

	fmt.Println(f)

	if err != nil {
		fmt.Println("Falha ao abrir o arquivo", err)
	}

	fmt.Println(">>>", s)

	// Tentando erroneamente alterar o tipo de uma variável previamente declarada de inteiro para string*/
	j := 0
	e, j := "foo", "bar"

	fmt.Println(j)
}
