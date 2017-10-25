# [Code Militia](/index.md)

## Copiando slices

A função builtin [copy()](https://golang.org/pkg/builtin/#copy) tem por intuito copiar os items de um slice origem (***org***) para um slice destino(***dst***)

#### https://golang.org/pkg/builtin/#copy
> The copy built-in function copies elements from a source slice into a destination slice. (As a special case, it also will copy bytes from a string to a slice of bytes.) The source and destination may overlap. Copy returns the number of elements copied, which will be the minimum of len(***src***) and len(***dst***).



> Afunção embutida ***copy***(), copia elementos de um slice origem para um slice destino.
(Em um caso especial, também copia bytes de uma string para um slice de bytes) A origem e o destino podem se sobrepor. Copy retorna a quantidade de elementos copiados, que será no minimo de len(***src***) e len(***dst***)


## Exemplo
```go 
package main

import "fmt"

func main() {
	var org = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var dst = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	numOfItems := copy(dst, org)

	fmt.Printf("%d items copiados.\n", numOfItems)

	for k, v := range dst {
		fmt.Printf("%d => %d\n", k, v)
	}
}
```
### Saída
```
10 items copiados.
0 => 0
1 => 1
2 => 2
3 => 3
4 => 4
5 => 5
6 => 6
7 => 7
8 => 8
9 => 9
```

## Copiando uma string

```go
package main

import "fmt"

func main() {
	org := "Olá mundo"
	var dst = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	copy(dst, org)

	fmt.Printf("%s \n", dst)
}
```

### Saída

```
Olá mundo
```

## Pequeno truque na função copy()

Um dos problemas que tive implementando essa função pela primeira vez, é que não sabia que a quantidade de elementos que serão copiados pela função copy é igual a quantidade de elemtos do slice ***dst***. Então:

```
numOfItemsCopied = len(dst)
```

### Exemplo:
```go
package main

import (
	"fmt"
)

func main() {
	var org []int = []int{1, 2, 3, 4, 5, 6}
	var dst []int

	// Copiando 
	copy(dst, org)

	// Let's print it !!! 
	for i := range s2 {
		fmt.Println(i)
	}
}
```

### :(
Nada foi impresso, não é mesmo? Isso porque ***dst*** está vazio.
copy() copia len(***dst***) elementos de ***org***, ou nesse caso, 0


```go
// Doing it right!!
package main


import (
	"fmt"
)

func main() {
    // Definindo um tamanho inicial para dst igual ao tamanho de org
	dst = make([]int, len(org))

	copy(dst, org)

    // Let's print it !!! 
	for k, i := range dst {
		fmt.Println(k, "=>", i)
	}
}
```

### Saída
```
0 => 1
1 => 2
2 => 3
3 => 4
4 => 5
5 => 6
```

Por hoje é só. &#128640;