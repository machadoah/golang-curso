package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     uint
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome     string
	Idade    uint
	Endereco // composicao em go
	Ativo    bool
}

func main() {
	tonho := Cliente{
		Nome:  "Antonio",
		Idade: 21,
		Ativo: true,
	}

	tonho.Logradouro = "Rua A"
	tonho.Numero = 1
	tonho.Cidade = "RJ"
	tonho.Estado = "RJ"

	wes := Cliente{
		Nome:  "Wesley",
		Idade: 9,
		Ativo: true,
	}

	wes.Logradouro = "Rua A"
	wes.Numero = 1
	wes.Cidade = "RJ"
	wes.Estado = "RJ"

	usuarios := []Cliente{tonho, wes}

	for indice, user := range usuarios {
		if user.Ativo {
			fmt.Printf("%v - Me chamo %v e tenho %v anos\n", indice, user.Nome, user.Idade)
			fmt.Printf("Eu moro na rua %v, numero %v que fica em %v no estado %v\n", user.Logradouro, user.Numero, user.Cidade, user.Estado)
			fmt.Println("--------------------------")
		}
	}
}
