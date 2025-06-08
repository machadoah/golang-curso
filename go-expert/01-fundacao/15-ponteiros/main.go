package main

func main() {
	// 1. Declaração e atribuição
	a := 10
	println("Valor de a:", a)     // valor é 10
	println("Endereço de a:", &a) // ex: 0x14000058730

	// 2. Referência a um endereço de memória (ponteiro)
	var ponteiroDeA *int = &a
	println("ponteiroDeA (endereço):", ponteiroDeA)

	/*
		3. Desreferenciação: o operador * pega o valor onde o ponteiro aponta.

		*ponteiroDeA = 20 modifica o valor em 'a', pois ponteiroDeA referencia o
		mesmo endereço de memória de a.
	*/
	*ponteiroDeA = 20
	println("Novo valor de a:", a) // 20

	// 4. Forma abreviada: inferência automática de ponteiro com :=
	b := &a
	println("b (endereço):", b)
	println("Valor apontado por b (*b):", *b)

	// 5. Modificando via b: altera o valor de a novamente
	*b = 90
	println("Novo valor de a após *b = 90:", a)

	// 🧠 Observação: acontece sem copiar valor, acessamos 'a' diretamente na memória
}
