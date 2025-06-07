package main

// declarando e atribuindo constante global
const a = "GO"

// declarando variavel booleana b
var b bool

// declarando variaveis
var (
	c int     // valor default 0
	d string  // valor default '' -> string vazia
	e float64 // valor default 0.0
)

var nome string = "Antonio"
var idade int = 21

// ponto de entrada da aplicacao
func main() {
	println(a) // GO
	println(b) // false

	// variavel local a
	var a string
	println(a) // '' -> espaco vazio

	/*
		nome = 1

		Algo como o acima, irá dar um erro:
			cannot use 1 (untyped int constant) as string value

		Respeitando a premissa que Go segue a risca os tipos definidos.
	*/

	nome = "a"

	/*
		var b = "x"

		retornara algo como:
			declared and not used

		em go é retornado erro de compilacao quando algo é declarado no escopo
		da funcao mas nao é usado
	*/

	// utilizando operador := para atribuicao e definicao implicita
	y := "y"
	println(y)

	y = "uai"
	println(y)
}
