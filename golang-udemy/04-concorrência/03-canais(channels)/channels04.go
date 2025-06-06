package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("OI", canal)

	fmt.Println("Depois da funcao escrever comecar a ser executada")

	// iterando o canal
	for msg := range canal {
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
