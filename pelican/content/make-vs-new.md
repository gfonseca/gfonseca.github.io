Title: Make VS New
Date: 2017-11-05 19:21
Category: golang

# Make() vs New()

Este post será breve, trata se de uma explicação das funções embutidas **make**() e **new**().

## New

A função embutida **[new](https://golang.org/pkg/builtin/#new)** aloca espaço na memória para um determinado tipo, e retorna um ponteiro para este endereço recém alocado

```go
var i *int
i = new(int)

fmt.Printf(“> %d”, *i)
```
Note que a nossa posição de memória já está inicializada com o seu respectivo *[zerovalue](https://tour.golang.org/basics/12)*, no caso 0, mas isto será abordado em outro post.

## Make

A função **[make](https://golang.org/pkg/builtin/#make)** funciona apenas para as primitivas slice, chanel e map, esses tipos em especial não podem ser apenas alocados, eles também precisam ser inicializados, com valores específicos.
Segundo o documento oficial [Effective GO](https://golang.org/doc/effective_go.html#allocation_make) estes tipos, são na verdade estruturas de dados que precisam ter seus campos preenchidos com os devidos valores. No caso do tipo slice por exemplo, trata se de uma estrutura com os campos, data (referência para um array), tamanho e capacidade. Diferente da função **new** a função **make** não retorna um ponteiro, mas sim um valor.

```go
var s []int = make([]int, 10)
fmt.Prinln(s)
```

## Declaração ilegal
```go
d := make(int)
```

**Saída**
```
code/make-vs-new.go:16: cannot make type int
```


## Finalizando &#128123;

Apesar da ambiguidade que os termos make e new podem gerar, essas duas funções tem propósitos totalmente diferente. A função **new** realiza alocação de memória de forma similar a função maloc da linguagem C, enquanto a função **make** tem o propósito de inicializar primitivas especiais da linguagem.

Até a próxima &#128640;






