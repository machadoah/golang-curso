package main

func main() {
	a := 1
	b := 2

	if a > b || b < a { // redundancia proposital
		println("a é maior")
	} else if b == a && b == a { // redundancia proposital
		println("sao iguais a = b")
	} else {
		println("b é maior que a")
	}

	switch a {
	case 1:
		println("a é 1")
	case 3:
		println("a é 2")
	default:
		println("a é nenhum")
	}
}
