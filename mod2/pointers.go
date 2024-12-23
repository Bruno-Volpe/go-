package main

import "fmt"

func a() {
	x := 1
	y := 2

	var ip *int = &x
	ip = &y

	nome := new(string)
	*nome = "bruno"

	fmt.Printf("O meu nome eh: %s\n", *nome)
	fmt.Printf("O valor de x é: %d\n", ip)
}

/*
	& -> retorna o endereço de memória
	* -> retorna o valor do endereço de memória

	new -> aloca memória para um tipo de dado e retorna um ponteiro para o tipo de dado alocado = uma nova forma de declarar variavel mas retorna o endereco na memmoria
*/
