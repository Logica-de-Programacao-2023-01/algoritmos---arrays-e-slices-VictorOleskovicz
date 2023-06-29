package main

import "fmt"

func main() {
	slice := make([]int, 0, 5)
	var num int

	fmt.Println("Digite um número inteiro:")

	fmt.Print("Número: ")
	fmt.Scan(&num)

	found := false

	for _, n := range slice {
		if n == num {
			found = true
			break
		}
	}

	if !found {
		slice = append(slice, num)
	}

	fmt.Println("\nSlice resultante:", slice)
}
