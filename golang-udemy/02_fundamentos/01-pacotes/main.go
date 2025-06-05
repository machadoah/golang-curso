package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"pac/utils"
)

func main() {
	fmt.Println("Hola, Golang!")
	utils.Escrever()

	erro := checkmail.ValidateFormat("machadoah@proton.me")
	fmt.Println(erro)
}
