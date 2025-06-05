package main

import "fmt"

func main() {
	// variavel := map[tipo_chaves]tipo_valores
	usuario := map[string]string{
		"nombre":      "Banana",
		"sobrenombre": "Nanica",
	}

	fmt.Println(usuario["nombre"], usuario["sobrenombre"]) // Banana Nanica
	fmt.Println(usuario)                                   // map[nombre:Banana sobrenombre:Nanica]

	// chabes string e valor é um map com chave string e valor string
	usuarioA := map[string]map[string]string{
		"nombre": {
			"primeiro": "Antonio",
			"ultimo":   "Banana",
		},
		"curso": {
			"nome":        "ADS",
			"instituicao": "fatec",
		},
	}

	fmt.Println(usuarioA["nombre"]["ultimo"]) // Banana
	fmt.Println(usuarioA)

	// apagando chaves
	delete(usuarioA, "nombre")
	fmt.Println(usuarioA)

	// adicionando chave
	usuarioA["nome"] = map[string]string{
		"primeiro": "Antonio",
		"ultimo":   "Machado",
	}

	fmt.Println(usuarioA)
}
