package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// cria wait group
	var waitGroup sync.WaitGroup

	// quantidade de goroutines
	waitGroup.Add(2)

	// funcao anonima com a 1a goroutine
	go func() {
		escrever("Github")
		waitGroup.Done() // -1
	}()

	// funcao anonima com a 2a goroutine
	go func() {
		escrever("Programando em Go")
		waitGroup.Done() // -1
	}()

	// Espera -> passa quando waitgroup for zero
	waitGroup.Wait() // 0
}

func escrever(texto string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
