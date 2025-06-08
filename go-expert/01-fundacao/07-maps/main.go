package main

import "fmt"

func main() {

	salarios := map[string]uint{
		"Wesley":  1000,
		"Antonio": 3000,
	}

	for name, salary := range salarios {

		fmt.Printf("O salário do %v é R$%v\n", name, salary)
	}

	// blank identifier
	for _, salary := range salarios {

		fmt.Printf("O salário é R$%v\n", salary)
	}

	delete(salarios, "Wesley")

	salarios["Wes"] = 6000

	fmt.Println(salarios["Wes"])
}
