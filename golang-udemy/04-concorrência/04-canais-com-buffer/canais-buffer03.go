package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "OI"
	canal <- "OLÁ"

	msgA := <-canal
	msgB := <-canal

	fmt.Println(msgA)
	fmt.Println(msgB)

}
