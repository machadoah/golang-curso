package main

import (
	"io"
	"net/http"
)

func main() {
	chamada, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}

	resultado, err := io.ReadAll(chamada.Body)
	if err != nil {
		panic(err)
	}

	println(string(resultado))

	chamada.Body.Close()
}
