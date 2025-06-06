package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Github") // go routine -> execute, mas nao espere terminar
	escrever("Programando em Go")

}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
