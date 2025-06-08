package main

import (
	"errors"
	"fmt"
)

// ponto de entrada do go
func main() {
	fmt.Println(sumV1(1, 2))
	fmt.Println(sumV2(1, 2))

	res, isEven := sumIsEven(2, 2)

	var resEvenOrOdd string

	if isEven {
		resEvenOrOdd = "Par"
	} else {
		resEvenOrOdd = "Impar"
	}

	fmt.Printf("2 mais 2 é %v e esse numero é %v\n", res, resEvenOrOdd)

	// tratamento de erro em go
	val, err := isEvenF(10, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}
}

func sumV1(a int, b int) int {
	return a + b
}

func sumV2(a, b int) int {
	return a + b
}

func sumIsEven(a, b int) (int, bool) {
	sum := a + b
	isEven := sum%2 == 0

	return sum, isEven
}

// funcoes com erros
func isEvenF(a, b int) (bool, error) {

	res := a + b
	if res%2 == 0 {
		return true, nil
	}

	return false, errors.New("O número não é PAR!")
}
