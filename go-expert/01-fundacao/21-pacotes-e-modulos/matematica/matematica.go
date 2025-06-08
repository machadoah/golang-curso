package matematica

import "math"

// Soma : exportavel		-> publico
// soma : nao exportavel	-> privado
func Soma[T int | float64](a ...T) T {
	var soma T

	for _, num := range a {
		soma += num
	}

	return soma
}

var Pii = math.Pi

type Carro struct {
	Marca string
	placa string // "atributo privado"
}
