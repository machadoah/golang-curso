package main

import "fmt"

func fib(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fib(posicao-2) + fib(posicao-1)
}

func worker(tasks chan uint, res chan uint) {
	// taks <-chan uint		- só envia dados
	// res chan<- uint		- só recebe dados
	for task := range tasks {
		res <- fib(task)
	}

}

func main() {
	tarefas := make(chan uint, 45)
	resultados := make(chan uint, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := uint(0); i < 45; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := uint(0); i < 45; i++ {
		res := <-resultados
		fmt.Println(res)
	}

}
