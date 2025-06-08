package main

import "fmt"

// todo mundo implementa essa interface
type x interface{} // interface{} can be replaced by any modernize(efaceany)

func main() {

	var y interface{} = 10
	var z any = "Olá"

	showType(y) // O tipo da variavel é int e o valor é 10
	showType(z) // O tipo da variavel é string e o valor é Olá

}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é %T e o valor é %v\n", t, t)
}
