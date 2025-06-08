package main

func soma(a, b *int) int {

	*a = 50
	*b = 50
	return *a + *b
}

func main() {
	minhaVar01 := 10
	minhaVar02 := 20

	res := soma(&minhaVar01, &minhaVar02)

	println(res)        // 100
	println(minhaVar01) // 50
	println(minhaVar02) // 50
}
