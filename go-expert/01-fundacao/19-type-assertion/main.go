package main

import "fmt"

func main() {
	var minhaVar any = "Antonio Henrique"

	println(minhaVar)          // (0x1050796a0,0x1050863e8)
	println(minhaVar.(string)) // Antonio Henrique

	if res, ok := minhaVar.(int); ok {
		println(res)
	} else {
		println("Não foi possivel converter")
	}

	// panic
	fmt.Println(minhaVar.(int)) // panic: interface conversion: interface {} is string, not int
}
