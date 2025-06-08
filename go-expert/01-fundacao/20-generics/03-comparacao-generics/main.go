package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Compara[T Number](a, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	println(Compara(10, 10.0))
}
