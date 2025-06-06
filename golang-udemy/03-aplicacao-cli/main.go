package main

import (
	"command-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("🏁 Ponto de partida")

	application := app.Gerar()
	err := application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
