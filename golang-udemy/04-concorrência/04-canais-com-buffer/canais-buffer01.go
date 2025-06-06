package main

import "fmt"

func main() {

	// caso de deadlock

	canal := make(chan string)

	// essa linha fica aguardando alguem receber esse calor, para poder passar para o canal
	canal <- "OI"

	// nunca vai chegar aqui!
	msg := <-canal

	fmt.Println(msg)
}
