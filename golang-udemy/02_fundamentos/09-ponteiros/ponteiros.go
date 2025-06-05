package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2) // 10 10

	variavel1++                       // adiciona 1 a variavel1
	fmt.Println(variavel1, variavel2) // 11 10

	// ponteiro é uma referencia de memoria
	var variavel3 int
	var ponteiro *int
	fmt.Println(variavel3, ponteiro) // 0 <nil>

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro)  // 100 0x14000098028
	fmt.Println(variavel3, *ponteiro) // desferenciacao -> 100 100

	variavel3 = 150
	fmt.Println(variavel3, ponteiro)  // 150 0x14000098028
	fmt.Println(variavel3, *ponteiro) // 150 150
}
