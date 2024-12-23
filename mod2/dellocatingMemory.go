package main

import "fmt"

func x() {
	x := "bruno"

	fmt.Printf("O valor de x é: %s\n", x)
}

/*
	Stack - coisas que ficam dentor do escopo de uma funcao e sao desalocas quando a funcao termina
	Heap - ficam forma do escopo da funcao e precisam de desalocadas - NAO em GO por conta do Garbage Collector
*/
