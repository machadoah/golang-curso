package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("OI", canal)

	fmt.Println("Depois da funcao escrever comecar a ser executada")

	// espera receber infinitas msgs
	for {
		msg := <-canal
		fmt.Println(msg)
	}

	// DEADLOCK!

	/*
		fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan receive]:
		main.main()

		---

		Quando nao existe mais nada que envia dados para o canal, mas o canal
		nao recebe mais dados.
	*/
}

func escrever(texto string, canal chan string) {
	// envia somente 5 msgs
	for i := 1; i <= 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
}
