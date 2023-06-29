package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite o tamanho dos arrays: ")
	fmt.Scan(&n)

	array1 := make([]int, n)
	array2 := make([]int, n)

	fmt.Println("Digite os elementos do primeiro array:")
	readArray(array1)

	fmt.Println("Digite os elementos do segundo array:")
	readArray(array2)

	sumArray := addArrays(array1, array2)

	fmt.Println("Array
