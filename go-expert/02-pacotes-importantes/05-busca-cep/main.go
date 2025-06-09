package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
	API: https://viacep.com.br/
*/

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, url := range os.Args[1:] { // go run ./main.go ola
		println("digitado =", url)

		// req, err := http.Get(url)
		req, err := http.Get("http://viacep.com.br/ws/" + url + "/json/")
		if err != nil {
			fmt.Printf("Erro ao fazer requisição: %v\n\t\t%v\n", url, err)
		}

		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Printf("Erro ao ler arquivo.\n\t\t%v\n", err)
		}

		var data ViaCEP
		if err := json.Unmarshal(res, &data); err != nil {
			fmt.Printf("Erro ao fazer unmarshal (parse das respostas).\n\t\t%v\n", err)
		}

		fmt.Println(data)

		file, err := os.Create("cep.txt")
		if err != nil {
			fmt.Printf("Erro ao criar arquivo.\n\t\t%v\n", err)
		}
		defer file.Close()

		if _, err := file.WriteString(fmt.Sprintf("CEP: %v - CIDADE: %v - ESTADO: %v\n", data.Cep, data.Localidade, data.Uf)); err != nil {
			fmt.Printf("Erro ao escrever no arquivo.\n\t\t%v\n", err)
		} else {
			fmt.Println("Foi escrito com sucesso no arquivo cep.txt")
		}

	}
}
