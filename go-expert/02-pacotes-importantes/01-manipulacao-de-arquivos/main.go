package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Criando arquivo
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// escrita em arquivos

	// tam, err := f.WriteString("OLÁ\n")
	tam, err := f.Write([]byte("Um slice de algum dado"))
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %v bytes\n", tam)

	f.Close()

	// leitura
	a, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(a))

	// leitura em partes
	a2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(a2) // criacao do reader
	buffer := make([]byte, 10)    // criacao do buffer

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

}
