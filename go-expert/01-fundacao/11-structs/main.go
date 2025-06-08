package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade uint
	Ativo bool
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
		Ativo: false,
	}

	usuarios := []Cliente{tonho, wes}

	for indice, user := range usuarios {
		if user.Ativo {
			fmt.Printf("%v - Me chamo %v e tenho %v anos\n", indice, user.Nome, user.Idade)
		}
	}
}
