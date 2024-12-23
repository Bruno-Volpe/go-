package main

import "fmt"

func main() {
	var x int32 = 1
	var y int64 = 2

	// x = y // error
	x = int32(y) // convertion

	fmt.Printf("O valor de x é: %d\n", x)
}
