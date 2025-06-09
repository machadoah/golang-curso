package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"-"`     // implementacao de tags
	Saldo  int `json:"saldo"` // - omite o valor no json
}

func main() {
	// criacao de um struct
	conta := Conta{Numero: 01, Saldo: 100}

	res, err := json.Marshal(conta) // Marshar é transformar em json!
	if err != nil {
		println(err)
	}

	println(string(res))

	// entrega para algo -> stdout aqui
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(conta)

	// JSON -> STRUCT
	var contaA Conta

	// jsonPuro := []byte(`{"Numero": 2, "Saldo": 50000}`)
	jsonPuro := []byte(`{"Numero": 2, "Saldo": 50000}`)

	if err := json.Unmarshal(jsonPuro, &contaA); err != nil {
		fmt.Println(err)
	}

	fmt.Println(contaA) // 2 50000
}
