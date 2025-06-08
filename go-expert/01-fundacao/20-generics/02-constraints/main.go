package main

type MyNumber int

// criando constraints para aceitar int e float64
type Number interface {
	~int | float64 // o til, permite todos os demais tipos que sao int
}

// uso de generics nos tipos recebidos
func Soma[T Number](m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}

	return soma
}

func main() {
	mi := map[string]int{
		"Antonio": 3000,
		"Joao":    1000,
		"Caio":    6200,
	}

	mf := map[string]float64{
		"Antonio": 3000.90,
		"Joao":    1000.65,
		"Caio":    6200.10,
	}

	mj := map[string]MyNumber{
		"Antonio": 3000,
		"Joao":    1000,
		"Caio":    6200,
	}

	println(Soma(mi))
	println(Soma(mf))
	println(Soma(mj))
}
