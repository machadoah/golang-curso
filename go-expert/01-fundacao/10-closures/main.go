package main

import "fmt"

func main() {

	// funcao interna anonima
	letraA, isA := func() (string, bool) {
		return returnA(), true
	}()

	fmt.Println(letraA, isA)

	funcA := returnFunc()

	funcA()
}

func returnA() string {
	return "A"
}

// funcao que recebe outra funcao
func returnFunc() func() {
	return func() {
		fmt.Println("EXECUTANDO A FUNCAO RETURNFUNC")
	}
}
