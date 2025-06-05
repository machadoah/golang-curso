package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

type boletim struct {
	melhorDisciplina string
	media            int
}

type aluno struct {
	nome    string
	idade   uint8
	boletim boletim
}

func main() {
	fmt.Println("Arquivo de Structs")

	var u usuario
	fmt.Println(u) // { 0} -> todos os campos com valores vazios

	u.nome = "Antonio"
	u.idade = 21

	fmt.Println(u) // {Antonio 21}

	user := usuario{"Antonio Henrique", 21}
	fmt.Println(user) // {Antonio Henrique 21}

	// construindo struct com somente um valor
	user1 := usuario{idade: 21}
	fmt.Println(user1) // { 21}

	// Struct que o campo é um struct
	boletim1 := boletim{"Matematica", 8}
	aluno1 := aluno{"Isa", 19, boletim1}

	fmt.Println(aluno1) // {Isa 19 {Matematica 8}}

}
