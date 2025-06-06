package main

import (
	"fmt"
	"time"
)

func main() {
	chan01, chan02 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Second / 2)
			chan01 <- "CANAAAL 01"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			chan02 <- "CANAAAL 02"
		}
	}()

	for {
		mensgaemC1 := <-chan01
		fmt.Println(mensgaemC1)

		mensgaemC2 := <-chan02
		fmt.Println(mensgaemC2)
	}

	/*
		CANAAAL 01

		-- espera 2 segundos

		CANAAAL 02
		CANAAAL 01

		-- espera 2 segundos

		CANAAAL 02
		CANAAAL 01
	*/
}
