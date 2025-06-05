package main

import "fmt"

func main() {
	fmt.Println("Arrays e slices")

	// [numero de posicoes]tipo de dados
	var arrayA [3]int
	fmt.Println(arrayA) // [0 0 0]

	// populando array
	arrayA[0] = 10
	arrayA[1] = 20
	arrayA[2] = 30
	fmt.Println(arrayA) //[10 20 30]

	arrayB := [3]string{"a", "b", "c"}
	fmt.Println(arrayB) //[a b c]

	// determinando tamanho do array conforme atribuicao
	arrayC := [...]string{"a", "b", "c"}
	fmt.Println(arrayC)

}
