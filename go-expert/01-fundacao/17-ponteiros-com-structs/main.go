package main

import "fmt"

type Cliente struct {
	nome    string
	posicao int
}

func NewCliente() *Cliente {
	return &Cliente{}
}

func (c *Cliente) andou() {
	c.posicao = c.posicao + 1

	fmt.Printf("O cliente %v andou! Sua posicao é %v\n", c.nome, c.posicao)
}

func main() {

	a := Cliente{
		nome:    "Antonio",
		posicao: 10,
	}

	a.andou()

	fmt.Printf("O valor da variavel nome é %v e posicao é %v\n", a.nome, a.posicao)
}
