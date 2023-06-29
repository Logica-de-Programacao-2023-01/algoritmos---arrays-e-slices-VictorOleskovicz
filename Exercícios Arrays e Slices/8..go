package main

import "fmt"

func main() {
	slice := []string{"apple", "banana", "orange", "apple", "grape", "apple", "kiwi", "apple"}
	var value string

	fmt.Println("Digite um valor:")

	fmt.Print("Valor: ")
	fmt.Scan(&value)

	result := make([]string, 0)

	for _, item := range slice {
		if item != value {
			result = append(result, item)
		}
	}

	fmt.Println("\nSlice resultante:", result)
}
