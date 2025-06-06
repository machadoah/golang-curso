package main

import "fmt"

type generic interface {
}

func generica(g generic) { // podemos usar any no lugar de generic
	fmt.Println(g)
}

func main() {
	generica("antonio")
	generica(1)
	generica(true)

	mapa := map[any]any{
		1:    "antonio",
		true: 9,
	}

	fmt.Println(mapa)

}
