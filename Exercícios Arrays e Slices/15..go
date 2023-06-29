package main

import "fmt"

func main() {
	array := [10]float64{3.2, 7.5, 9.8, 1.2, 6.5, 4.7, 8.9, 2.3, 5.1, 10.4}

	slice := make([]float64, 0)

	for _, num := range array {
		if num > 5 {
			slice = append(slice, num)
		}
	}

	fmt.Println("Novo Slice:", slice)
}
