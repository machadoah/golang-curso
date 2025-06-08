package main

func SomaInteiro(m map[string]int) int {
	var soma int

	for _, v := range m {
		soma += v
	}

	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64

	for _, v := range m {
		soma += v
	}

	return soma
}

// uso de generics nos tipos recebidos
func Soma[T int | float64](m map[string]T) T {
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

	println(SomaInteiro(mi))

	mf := map[string]float64{
		"Antonio": 3000.90,
		"Joao":    1000.65,
		"Caio":    6200.10,
	}

	println(SomaFloat(mf))

	println("----------------------")

	println(Soma(mi))
	println(Soma(mf))
}
