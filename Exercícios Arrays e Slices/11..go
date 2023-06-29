package main

import "fmt"

func main() {
	var matrix [2][3]int

	fmt.Println("Informe os valores da matriz:")

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Matriz[%d][%d]: ", i, j)
			fmt.Scan(&matrix[i][j])
		}
	}

	var row, col int
	fmt.Print("Informe o índice da linha: ")
	fmt.Scan(&row)
	fmt.Print("Informe o índice da coluna: ")
	fmt.Scan(&col)

	fmt.Printf("\nValor na posição [%d][%d]: %d\n", row, col, matrix[row][col])
}
