package main

import (
	"fmt"
	"time"
)

func main() {
	// for convencional
	i := 0

	for i < 10 {
		i++
		fmt.Println("i =", i)
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	// Com com iterador/Contador
	for j := 0; j <= 10; j++ {
		fmt.Println("j =", j)
		time.Sleep(time.Second)
	}

	// For com Range
	nomes := []string{"Joao", "Jose", "Davi"}

	for i, nome := range nomes {
		fmt.Println(i, " - nome =", nome)
	}

	// For ignorando o indice com range
	for _, nome := range nomes {
		fmt.Println("nome =", nome)
	}

	// Iterando uma string
	for _, letra := range "ANTONIO" {
		fmt.Println(string(letra)) // se nao fazer o cast vai trazer numeros da ascii
	}

	// Iterando um map
	usuarios := map[string]int{"Joao": 1, "Jose": 2, "Davi": 3}

	for key, value := range usuarios {
		fmt.Println(key, value*10)
	}

	//Loop infinito
	for {
		fmt.Println("Para sempre ...")
	}
}
