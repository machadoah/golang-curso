package main

import (
	"fmt"
	"time"
)

func main() {
	// 1) aqui é criado um canal de string, sendo assim recebe e envia strings
	canal := make(chan string)

	// chama a funcao como goroutine e passa o canal
	go escrever("OI", canal)

	fmt.Println("Depois da funcao escrever comecar a ser executada")

	// canal espera o valor e passa para msg
	msg := <-canal
	fmt.Println(msg)
}

// criacao da funcao escrever que recebe a string e o canal
func escrever(texto string, canal chan string) {
	for i := 1; i <= 5; i++ {
		// o canal ira receber o texto - passa o valor pelo canal
		canal <- texto
		time.Sleep(time.Second)
	}
}
