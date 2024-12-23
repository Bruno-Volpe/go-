package main

import "fmt"

func foo() *int {
	x := 1
	return &x
}

func qq() {
	y := foo()
	fmt.Println(*y)
}

/*
Aqui, mesmo quando foon termina de retronar a variavel x nao eh excluida, posi pode ser
que main use essa variavel pois ele possui um ponteiro para ela

Para isso serve a gabrbage collector em GO
ele so libera quando todos os ponteiros ja foram usados
*/
