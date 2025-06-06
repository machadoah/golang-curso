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
		// o 2o parametro é se o canal esta aberto ou nao
		msg, open := <-canal

		// se ele estiver fechado sai do for
		if !open {
			break // sai de um loop
		}

		fmt.Println(msg)
	}

	fmt.Println("🏁 fim do programa")
}

func escrever(texto string, canal chan string) {
	// envia somente 5 msgs
	for i := 1; i <= 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	// fecha o canal apos a 5 iteracoes
	close(canal)
}
