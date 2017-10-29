# Reatribuição de variáveis

Saudações a todos. &#128075;
Hoje gostaria de falar sobre um pequeno truque envolvendo reatribuição de variáveis que pode ser desconhecidos pelos neófitos da linguagem, e que acredito, pode ser muito relevante no desenvolvimento de algoritmos.

Vamos direto ao código.

```go
f, err := os.Open(name)
d, err := f.Stat()
```

Esta é uma operação legal!

Na maioria dos casos quando utilizamos a [Short variable delcarations](https://golang.org/ref/spec#Short_variable_declarations) ( := ) o compilador não nos permite utilizá lo uma segunda vez com a mesma variável, já que essa operação nada mais é do que uma abreviação de uma declaração tradicional de variável.

```go
var d int = 1;
// equivale a:
var d := 1
```

### Vejamos um exemplo

```go
var d int = 1;
d := 1;
```

### Saída
```
./reatribuicao-de-variaveis.go:9: no new variables on left side of :=
```

Como eu havia dito, o compilador não permite redeclarar uma variável no mesmo escopo, porém, por pragmatismo, existe uma situação em que você pode utilizar uma variável previamente declarada com o *Short variable declaration* e o compilador já havia nos dado está dica:<br/><br/>

**no new variables on left side of :=**
<br/><br/>


### Resolvendo o problema

https://golang.org/doc/effective_go.html#redeclaration
>In a := declaration a variable v may appear even if it has already been declared, provided:

>- this declaration is in the same scope as the existing declaration of v (if v is already declared in an outer scope, the declaration will create a new variable §),
>- the corresponding value in the initialization is assignable to v, and
>- there is at least one other variable in the declaration that is being declared anew.

A solução é simples, caso exista ao menos uma variável ainda não declarada do lado esquerdo do operador := é permitido utilizá lo para alterar valores de variáveis já declaradas.

```go
package main

import (
	"fmt"
	"os"
)

func main() {

    f, err := os.Open("/tmp/reasign")
    s, err := f.Stat()

    fmt.Println(f)

    if err != nil {
        fmt.Println("Falha ao abrir o arquivo", err)
    }

    fmt.Println(">>>", s)
}
```

A existência da variável ***d*** no comando ***d, err := f.Stat()*** permitiu a reatribuição da variável ***err***.

Perceba que isso não passa de uma reatribuição, e não uma redeclaração de variáveis, em outras palavras, este atalho não nos permite alterar o tipo de uma variável previamente declarada, o que torna a operação  a seguir ilegal.

```go
    j := 0
    e, j := "foo", "bar"
```

### Saída
```
cannot use "bar" (type string) as type int in assignment
```

### Finalizando &#127859;
Este conhecimento vem a ser útil pois pelo que tenho visto até o momento, o uso do nome de variável *err* é muito comum no tratamento de exceções, já que não temos exceções nativas em golang e o retorno de múltiplos valores por funções se torna o principal dispositivo na captura de falhas.

```go
f, err := os.Open("/tmp/reasign")
s, err := f.Stat()

fmt.Println(f)

if err != nil {
    fmt.Println("Falha ao abrir o arquivo", err)
}

fmt.Println(">>>", s)
```

Por hoje é só. &#128640;