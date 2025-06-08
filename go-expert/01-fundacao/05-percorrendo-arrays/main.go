package main

import "fmt"

func main() {
	/*
		Array em go é um estrura com tamanho definido
	*/

	// crindo array de inteiros com 3 posicoes
	var meuArray [3]int

	// preenchendo campos do array
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	// ultimo valor do array
	fmt.Println(meuArray[len(meuArray)-1]) // 30

	for i, v := range meuArray {
		fmt.Printf("O valor do indice %v é %v\n", i, v)
	}
}
