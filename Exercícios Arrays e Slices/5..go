package main

import "fmt"

func main() {
	var matrix [3][2]int

	fmt.Println("Informe os valores da matriz:")

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("Matriz[%d][%d]: ", i, j)
			fmt.Scan(&matrix[i][j])
		}
	}

	fmt.Println("\nMatriz resultante:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}
