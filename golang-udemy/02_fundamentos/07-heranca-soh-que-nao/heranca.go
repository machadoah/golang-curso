package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float64
}

type estudante struct {
	pessoa    // heranca em go, pega todos os campos de pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Heranca em Go")

	var p pessoa = pessoa{nome: "Antonio", sobrenome: "Machado", idade: 21, altura: 1.76}
	fmt.Println(p)

	var e estudante = estudante{pessoa: p, curso: "ADS", faculdade: "Fatec PG"}
	fmt.Println(e)

	fmt.Println("O nome do estudante é", e.nome) // Posso acessar estudante.atributo_de_pessoa
}
