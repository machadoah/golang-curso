package main

import "fmt"

func main() {
	sliceA := []int{1, 2, 3}
	fmt.Println(len(sliceA)) // 3 -> é a quantidade itens preenchidos
	fmt.Println(cap(sliceA)) // 3 -> capacidade de ocupacao do slice
	fmt.Println(sliceA)      // dados do slice -> [1 2 3]

	sliceB := append(sliceA, 4, 5) // crio um sliceB adicionando 4 e 5 no sliceA
	fmt.Println(len(sliceB))       // 5 -> quantidade de itens preenchidos
	fmt.Println(cap(sliceB))       // 6 -> capacidade - o go pega a quantidade i e multiplica por a 2
	fmt.Println(sliceB)            // [1 2 3 4 5]

	array := [...]int{1, 2, 3}
	sliceC := array[1:3]
	fmt.Println(sliceC) // [2 3]

	array[1] = 3
	fmt.Println(sliceC) // [3 3]
}
