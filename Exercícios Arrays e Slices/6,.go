package main

import "fmt"

func main() {
	array := [10]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	var value int

	fmt.Print("Digite um valor: ")
	fmt.Scan(&value)

	found := false

	for _, num := range array {
		if num == value {
			found = true
			break
		}
	}

	if found {
		fmt.Println("O valor está presente no Array.")
	} else {
		fmt.Println("O valor não está presente no Array.")
	}
}
