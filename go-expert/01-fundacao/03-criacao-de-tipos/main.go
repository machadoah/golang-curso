package main

type ID int

var (
	registroAluno ID     = 1
	nomeAluno     string = "Antonio"
)

func main() {
	println(registroAluno) // 1
	println(nomeAluno)     // Antonio
}
