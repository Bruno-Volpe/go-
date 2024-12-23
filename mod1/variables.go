package main

import "fmt"

type IDNum int

func teste() {
	var x int = 10
	var nome string = "Bruno"
	var preco float64 = 10.45
	var id IDNum = 123

	// Infering type
	var y = 20

	// lazy
	var d int
	d = 10

	// short declaration
	z := 30

	fmt.Printf("O valor de z é: %d\n", z)
	fmt.Printf("O valor de x é: %d\n", x)
	fmt.Printf("O nome é: %s\n", nome)
	fmt.Printf("O preço é: %.2f\n", preco)
	fmt.Printf("O ID é: %d\n", id)
	fmt.Printf("O valor de d é: %d\n", d)
	fmt.Printf("O valor de y é: %d\n", y)
}
