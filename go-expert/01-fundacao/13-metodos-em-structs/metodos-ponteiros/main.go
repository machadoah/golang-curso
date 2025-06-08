package main

import (
	"fmt"
)

type Cliente struct {
	Nome  string
	Idade uint
	Ativo bool
}

func (c *Cliente) Desativar() {
	c.Ativo = false

	fmt.Printf("o cliente %v foi destivado\n", c.Nome)
}

func ImprimeClientesAtivos(usuarios []*Cliente) {
	for indice, user := range usuarios {
		if user.Ativo {
			fmt.Printf("%v - Me chamo %v e tenho %v anos\n", indice, user.Nome, user.Idade)
		}
	}
}

func main() {
	tonho := Cliente{
		Nome:  "Antonio",
		Idade: 21,
		Ativo: true,
	}

	wes := Cliente{
		Nome:  "Wesley",
		Idade: 9,
		Ativo: true,
	}

	usuarios := []*Cliente{&tonho, &wes}

	ImprimeClientesAtivos(usuarios)

	tonho.Desativar()

	ImprimeClientesAtivos(usuarios)
}
