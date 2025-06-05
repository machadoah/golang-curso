package main

import "fmt"

var n int

func init() {
	fmt.Println("Aqui ela é executada sempre quando o arquivo é executado")
	n = 10
}

func main() {
	fmt.Println("------ FUNCAO MAIN ------")

	fmt.Println(n)
}
