package main

import "fmt"

func main() {
	canal := make(chan string, 2) // passa a capacidade para o canal - buffer

	canal <- "OI"

	msg := <-canal

	fmt.Println(msg)

	/*
		nao vai dar deadlock pois estou passando um item para ele e ele so trava
		quando ultrapassar a capacidade
	*/
}
