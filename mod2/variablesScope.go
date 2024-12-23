package main

import "fmt"

var resultado int

func sum(x int, y int) int {
	resultado := x + y // redeclara a variavel
	return resultado
}

func multiply(x int, y int) int {
	resultado = x * y // Subsituti e  resultado
	return resultado
}

func b() {
	fmt.Printf("O valor da soma é: %d\n", sum(2, 3))
	fmt.Printf("O valor da multiplicação é: %d\n", multiply(2, 3))
	fmt.Printf("O valor da variável resultado é: %d\n", resultado)
}
