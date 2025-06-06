package main

import "fmt"

var db = []usuario{}

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("log: estou no método salvar ... ")
	db = append(db, u)

	fmt.Printf("Usuário '%s' - salvos no db\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	fmt.Println("ESTADO ATUAL DO DB:", db)

	u01 := usuario{"AAbrao", 20}
	u01.salvar()

	fmt.Println("ESTADO ATUAL DO DB:", db)

	u02 := usuario{"Bernardo", 17}
	u02.salvar()

	fmt.Println("ESTADO ATUAL DO DB:", db)

	fmt.Println(u02.maiorDeIdade())

	fmt.Println(u02)
	u02.fazerAniversario()
	fmt.Println(u02)
}
