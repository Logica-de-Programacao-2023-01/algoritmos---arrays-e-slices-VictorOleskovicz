package main

import (
	"fmt"
	"math"
)

func main() {
	slice := []int{5, 2, 9, 1, 7, 4, 6, 3, 8, 10}

	min := math.MaxInt64
	max := math.MinInt64

	for _, num := range slice {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	fmt.Println("Valor mínimo:", min)
	fmt.Println("Valor máximo:", max)
}
