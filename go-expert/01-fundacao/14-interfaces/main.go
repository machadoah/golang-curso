package main

import (
	"fmt"
)

// qualquer struct que implementa Desativar() implementa Pessoa
type Pessoa interface {
	// os structs que implementam devem ter a mesma assinatura
	Saudar() string
}

type Cliente struct {
	Nome  string
	Idade uint
	Ativo bool
}

func (c Cliente) Saudar() string {
	return fmt.Sprintf("Olá sou um cliente e meu nome é %v", c.Nome)
}

type Empresa struct {
	Nome string
	Ramo string
}

func (e Empresa) Saudar() string {
	return fmt.Sprintf("Olá, sou uma empresa do ramo %v e meu nome é %v", e.Ramo, e.Nome)
}

func Saudando(p Pessoa) {
	fmt.Println(p.Saudar())
}

func main() {
	tonho := Cliente{
		Nome:  "Antonio",
		Idade: 21,
		Ativo: true,
	}

	wes := Empresa{
		Nome: "Wesley LTDA",
		Ramo: "Alimenticio",
	}

	Saudando(tonho)
	Saudando(wes)

}
