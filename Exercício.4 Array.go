package main

import "fmt"

func main() {
	var soma float32
	numeros := []float32{2, 3, 4}
	fmt.Println(numeros)

	for _, x := range numeros {

		soma = soma + x
	}
	media := soma / 6
	fmt.Print("A média é", media)

}
