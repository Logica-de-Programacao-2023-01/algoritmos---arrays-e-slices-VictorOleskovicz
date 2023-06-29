package main

import "fmt"

func main() {
	array := [10]int{2, 5, 3, 7, 1, 9, 4, 6, 8, 10}

	sum := 0

	for i, num := range array {
		if i%2 == 0 {
			sum += num
		}
	}

	fmt.Println("Soma dos elementos nas posições pares:", sum)
}
