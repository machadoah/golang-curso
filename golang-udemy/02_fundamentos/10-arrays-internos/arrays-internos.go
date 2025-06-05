package main

import "fmt"

func main() {
	// arrays internos
	slice := make([]float64, 10, 11) // tipo, tamanho, capacidade
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))

	slice = append(slice, 1.0)
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))

	slice = append(slice, 2.0)
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice)) // dobra o tamanho do slice!

	// slice com make sem capacidade, cap fica = len
}
